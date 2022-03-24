package controller

import (
	"majoo-backend-test/constant"
	"majoo-backend-test/helper"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (c *Controller) MerchantOutletOmzet(ctx echo.Context) error {
	response := new(constant.Response)

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*constant.JwtCustomClaims)
	actorID := claims.ID

	merchantIDStr := ctx.Param("merchant_id")
	outletIDStr := ctx.Param("outlet_id")

	startDateStr := ctx.QueryParam("start_date")
	endDateStr := ctx.QueryParam("end_date")
	pageStr := ctx.QueryParam("page")
	countStr := ctx.QueryParam("count")

	var merchantID, outletID int
	var err error

	if len(outletIDStr) > 0 {

		outletID, err = strconv.Atoi(outletIDStr)
		if err != nil {
			response.Message = constant.MSG_ERROR_INVALID_OUTLET_ID
			return ctx.JSON(http.StatusBadRequest, response)
		}

	} else if len(merchantIDStr) > 0 {

		merchantID, err = strconv.Atoi(merchantIDStr)
		if err != nil {
			response.Message = constant.MSG_ERROR_INVALID_MERCHANT_ID
			return ctx.JSON(http.StatusBadRequest, response)
		}

	} else {

		response.Message = constant.MSG_ERROR_INVALID_DATA
		return ctx.JSON(http.StatusBadRequest, response)

	}

	startDate, endDate, err := helper.ValidateStartEndDate(startDateStr, endDateStr)
	if err != nil {
		response.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, response)
	}

	page, count, err := helper.ValidatePagination(pageStr, countStr)
	if err != nil {
		response.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, response)
	}

	filters := constant.M{
		"start_date": startDate,
		"end_date":   endDate,
		"page":       page,
		"count":      count,
	}

	result, status, err := c.service.MerchantOutletOmzet(actorID, merchantID, outletID, filters)
	if err != nil {
		response.Message = err.Error()
		return ctx.JSON(status, response)
	}

	response.Data = result

	return ctx.JSON(http.StatusOK, response)
}
