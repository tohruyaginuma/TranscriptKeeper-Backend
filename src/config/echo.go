package config

import "github.com/labstack/echo/v4"

func SetEcho() *echo.Echo {
	e := echo.New()

	e.Debug = true
	e.HideBanner = false

	return e
}
