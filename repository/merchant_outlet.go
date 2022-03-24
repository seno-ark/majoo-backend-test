package repository

import (
	"log"
	"majoo-backend-test/model"
)

func (r *Repository) GetMerchant(merchantID int) (*model.MerchantOutlet, error) {

	var merchantOutlet = &model.MerchantOutlet{}

	query := "SELECT id AS merchant_id, user_id, merchant_name FROM Merchants WHERE id = ?"
	err := r.DB.Get(merchantOutlet, query, merchantID)
	if err != nil {
		log.Println(err.Error())
		return merchantOutlet, err
	}

	return merchantOutlet, nil
}

func (r *Repository) GetOutlet(outletID int) (*model.MerchantOutlet, error) {

	var merchantOutlet = &model.MerchantOutlet{}

	query := "SELECT m.user_id, m.id AS merchant_id, o.id as outlet_id, m.merchant_name, o.outlet_name FROM Merchants m INNER JOIN Outlets o ON o.merchant_id = m.id AND o.id = ?"
	err := r.DB.Get(merchantOutlet, query, outletID)
	if err != nil {
		log.Println(err.Error())
		return merchantOutlet, err
	}

	return merchantOutlet, nil
}
