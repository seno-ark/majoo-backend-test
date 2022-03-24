package controller

import (
	"net/http"

	"majoo-backend-test/constant"

	"github.com/labstack/echo/v4"
)

func (c *Controller) Login(ctx echo.Context) error {
	response := new(constant.Response)

	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if len(username) == 0 || len(password) == 0 {
		response.Message = constant.MSG_ERROR_INVALID_USERNAME_PASSWORD
		return ctx.JSON(http.StatusBadRequest, response)
	}

	result, status, err := c.service.Login(username, password)
	if err != nil {
		response.Message = err.Error()
		return ctx.JSON(status, response)
	}

	response.Message = "Success"
	response.Data = result

	return ctx.JSON(http.StatusOK, response)
}
