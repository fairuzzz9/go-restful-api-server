package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho"
	"gitlab.com/pos_malaysia/golib/database"
	"gitlab.com/pos_malaysia/golib/logs"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
)

var sqlStatements = map[string]string{
	// SQL Name : SQL statement
	// ok to have duplicate SQL statements, but not ok to have duplicate SQL names
	// Using map ensure that no duplicate keys. The compiler will stop if there's any duplicate in map literal

	//"GetAllFromTable": "SELECT * FROM SOMETABLE",

}

func init() {
	database.InitSQLStatements(sqlStatements)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func setupEcho() (*echo.Echo, zerolog.Logger) {

	// Setup Echo to use our logger
	logsConfig := logs.ConfigSet{}
	logger := logs.Configure(logsConfig)

	e := echo.New()
	e.Logger = lecho.New(logger) // Echo adapter for Zerolog

	e.Validator = &CustomValidator{validator: validator.New()}

	// Setup Echo's middleware
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))
	// log every request
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	return e, logger
}
