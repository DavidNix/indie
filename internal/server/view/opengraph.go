package view

import (
	"github.com/labstack/echo/v4"
)

const siteInfoKey = "site-info"

type SiteInfo struct {
	Name string
	Host string
}

func SetSiteInfo(info SiteInfo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(siteInfoKey, info)
			return next(c)
		}
	}
}

func GetSiteInfo(c echo.Context) SiteInfo {
	return c.Get(siteInfoKey).(SiteInfo)
}
