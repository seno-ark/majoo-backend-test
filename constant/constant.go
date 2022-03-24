package constant

import (
	"github.com/golang-jwt/jwt"
)

type M map[string]interface{}

type Response struct {
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
}

type JwtCustomClaims struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

const (
	CONFIG_MAX_PAGINATION_COUNT         = 20
	DATE_FILTER_FORMAT                  = "2006-01-02"
	MSG_ERROR_DATABASE                  = "Database Error"
	MSG_ERROR_TOKEN                     = "Token Error"
	MSG_ERROR_USER_NOT_FOUND            = "User not found"
	MSG_ERROR_MERCHANT_NOT_FOUND        = "Merchant not found"
	MSG_ERROR_OUTLET_NOT_FOUND          = "Outlet not found"
	MSG_ERROR_INVALID_USERNAME_PASSWORD = "Invalid username or password"
	MSG_ERROR_FORBIDDEN_REPORT          = "You don't have permission to access this report"
	MSG_ERROR_INVALID_DATA              = "Invalid Data"
	MSG_ERROR_INVALID_MERCHANT_ID       = "Invalid Merchant ID"
	MSG_ERROR_INVALID_OUTLET_ID         = "Invalid Outlet ID"
)
