package controllers

import (
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
    e := echo.New()

    e.GET("/health", healthCheckHandler)

    return e
}
