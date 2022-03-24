package model

import (
	"majoo-backend-test/constant"
	"time"
)

type OmzetByTime struct {
	Date  time.Time `db:"date" json:"date"`
	Total int64     `db:"total" json:"total"`
}

type OmzetReport struct {
	MerchantID   int            `json:"merchant_id"`
	OutletID     int            `json:"outlet_id"`
	MerchantName string         `json:"merchant_name"`
	OutletName   string         `json:"outlet_name"`
	Omzets       []*OmzetByTime `json:"omzets"`
	TotalData    int            `json:"total_data"`
	Filters      constant.M     `json:"filters"`
}
