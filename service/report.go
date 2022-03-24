package service

import (
	"database/sql"
	"errors"
	"majoo-backend-test/constant"
	"majoo-backend-test/model"
	"net/http"
	"time"
)

func (s *Service) MerchantOutletOmzet(actorID, merchantID, outletID int, filters constant.M) (*model.OmzetReport, int, error) {

	var merchantOutletDetail *model.MerchantOutlet
	var err error

	if outletID > 0 {

		merchantOutletDetail, err = s.repository.GetOutlet(outletID)
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New(constant.MSG_ERROR_OUTLET_NOT_FOUND)
		}
		if err != nil {
			return nil, http.StatusInternalServerError, errors.New(constant.MSG_ERROR_DATABASE)
		}

	} else {

		merchantOutletDetail, err = s.repository.GetMerchant(merchantID)
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New(constant.MSG_ERROR_MERCHANT_NOT_FOUND)
		}
		if err != nil {
			return nil, http.StatusInternalServerError, errors.New(constant.MSG_ERROR_DATABASE)
		}

	}

	if actorID != merchantOutletDetail.UserID {
		return nil, http.StatusForbidden, errors.New(constant.MSG_ERROR_FORBIDDEN_REPORT)
	}

	var startDate, endDate time.Time
	var page, count int

	if filters["start_date"] != nil {
		startDate = filters["start_date"].(time.Time)
	}
	if filters["end_date"] != nil {
		endDate = filters["end_date"].(time.Time)
	}
	if filters["page"] != nil {
		page = filters["page"].(int)
	}
	if filters["count"] != nil {
		count = filters["count"].(int)
	}

	timeDay := time.Hour * 24

	queryStartDate := startDate.Add(timeDay * time.Duration((page-1)*count))
	queryEndDate := queryStartDate.Add(timeDay * time.Duration(count))

	if queryStartDate.After(endDate.Add(time.Second * -1)) {
		return nil, http.StatusBadRequest, errors.New(constant.MSG_ERROR_INVALID_DATA)
	}
	if queryEndDate.After(endDate) {
		queryEndDate = endDate
	}

	totalDay := int(endDate.Sub(startDate).Hours() / 24)

	queryStartDateStr := queryStartDate.Format(constant.DATE_FILTER_FORMAT)
	queryEndDateStr := queryEndDate.Format(constant.DATE_FILTER_FORMAT)

	queryFilters := constant.M{
		"start_date": queryStartDateStr,
		"end_date":   queryEndDateStr,
	}
	results, err := s.repository.GetMerchantOutletOmzet(merchantID, outletID, queryFilters)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New(constant.MSG_ERROR_DATABASE)
	}

	mapResults := make(map[time.Time]*model.OmzetByTime)
	for _, v := range results {
		mapResults[v.Date] = v
	}

	newResults := []*model.OmzetByTime{}

	for i := 0; i < totalDay; i++ {
		newDate := queryStartDate.Add(timeDay * time.Duration(i))
		if i >= count || newDate.Add(time.Second).After(queryEndDate) {
			break
		}

		dailyOmzet := &model.OmzetByTime{
			Date: newDate,
		}
		if mapResults[newDate] != nil {
			dailyOmzet.Total = mapResults[newDate].Total
		}

		newResults = append(newResults, dailyOmzet)
	}

	omzetReport := &model.OmzetReport{
		Omzets:       newResults,
		MerchantID:   merchantOutletDetail.MerchantID,
		MerchantName: merchantOutletDetail.MerchantName,
		OutletID:     merchantOutletDetail.OutletID,
		OutletName:   merchantOutletDetail.OutletName,
		TotalData:    totalDay,
		Filters: constant.M{
			"page":  page,
			"count": count,
		},
	}

	return omzetReport, 0, nil
}
