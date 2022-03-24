package main

import (
	"log"
	"majoo-backend-test/constant"
	"majoo-backend-test/controller"
	"majoo-backend-test/helper"
	"majoo-backend-test/repository"
	"majoo-backend-test/service"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := helper.InitMySQL()
	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Ok.")
	})

	e.POST("/login", controller.Login)

	reportRoute := e.Group("/report")

	jwtConfig := middleware.JWTConfig{
		Claims:     &constant.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("TOKEN_SECRET")),
	}

	reportRoute.Use(middleware.JWTWithConfig(jwtConfig))

	reportRoute.GET("/merchant/:merchant_id/omzet", controller.MerchantOutletOmzet)
	reportRoute.GET("/outlet/:outlet_id/omzet", controller.MerchantOutletOmzet)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
