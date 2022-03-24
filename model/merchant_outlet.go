package model

type MerchantOutlet struct {
	UserID       int    `db:"user_id" json:"user_id"`
	MerchantID   int    `db:"merchant_id" json:"merchant_id"`
	MerchantName string `db:"merchant_name" json:"merchant_name"`
	OutletID     int    `db:"outlet_id" json:"outlet_id"`
	OutletName   string `db:"outlet_name" json:"outlet_name"`
}
