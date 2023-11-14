package server

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const baseURL = "https://CHANGEME.example.com"

const robotsTxt = `User-agent: *
Allow: /
Sitemap: %s/sitemap.xml
`

func robotsHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf(robotsTxt, baseURL))
}

func sitemapHandler(c echo.Context) error { return c.XML(http.StatusOK, nil) }
