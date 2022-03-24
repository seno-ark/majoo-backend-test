package helper

import (
	"errors"
	"majoo-backend-test/constant"
	"strconv"
	"time"
)

func ValidatePagination(pageStr, countStr string) (page, count int, err error) {

	if len(pageStr) > 0 {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			err = errors.New(constant.MSG_ERROR_INVALID_DATA)
			return
		}
	}

	if len(countStr) > 0 {
		count, err = strconv.Atoi(countStr)
		if err != nil {
			err = errors.New(constant.MSG_ERROR_INVALID_DATA)
			return
		}
	}

	if page < 1 {
		page = 1
	}
	if count < 1 || count > constant.CONFIG_MAX_PAGINATION_COUNT {
		count = constant.CONFIG_MAX_PAGINATION_COUNT
	}

	return
}

func ValidateStartEndDate(startDateStr, endDateStr string) (startDate, endDate time.Time, err error) {

	if len(startDateStr) < 1 {
		startDateStr = time.Now().Format(constant.DATE_FILTER_FORMAT)
	}
	if len(endDateStr) < 1 {
		endDateStr = time.Now().Format(constant.DATE_FILTER_FORMAT)
	}

	startDate, err = time.Parse(constant.DATE_FILTER_FORMAT, startDateStr)
	if err != nil {
		err = errors.New(constant.MSG_ERROR_INVALID_DATA)
		return
	}

	endDate, err = time.Parse(constant.DATE_FILTER_FORMAT, endDateStr)
	if err != nil {
		err = errors.New(constant.MSG_ERROR_INVALID_DATA)
		return
	}

	endDate = endDate.Add(time.Hour * 24)

	return
}
