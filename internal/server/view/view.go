package view

import "github.com/labstack/echo/v4"

func siteName(c echo.Context) string {
	return c.Get(siteName())
}
