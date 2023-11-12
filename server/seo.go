package server

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: Change the sitemap URL.
//go:embed views/robots.txt
var robotstxt string

func robotsHandler(c echo.Context) error { return c.String(http.StatusOK, robotstxt) }

func sitemapHandler(c echo.Context) error { return c.String(http.StatusOK, "TODO: Sitemap") }
