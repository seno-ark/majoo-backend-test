package repository

import (
	"log"
	"majoo-backend-test/constant"
	"majoo-backend-test/model"
)

func (r *Repository) GetMerchantOutletOmzet(merchantID, outletID int, filters constant.M) ([]*model.OmzetByTime, error) {

	var merchantOmzets = []*model.OmzetByTime{}

	var whereQuery string
	var args []interface{}

	if outletID > 0 {
		whereQuery = "outlet_id = ?"
		args = append(args, outletID)
	} else {
		whereQuery = "merchant_id = ?"
		args = append(args, merchantID)
	}

	args = append(args, filters["start_date"])
	args = append(args, filters["end_date"])

	query := "SELECT DATE(created_at) AS date, sum(bill_total) AS total FROM Transactions WHERE " + whereQuery + " AND DATE(created_at) >= ? AND DATE(created_at) < ? GROUP BY DATE(created_at) ORDER BY DATE(created_at) ASC"
	err := r.DB.Select(&merchantOmzets, query, args...)
	if err != nil {
		log.Println(err.Error())
		return merchantOmzets, err
	}

	return merchantOmzets, nil
}
