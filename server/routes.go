package server

import (
	"net/http"

	"github.com/DavidNix/indie/ent"
	"github.com/labstack/echo/v4"
)

func addRoutes(app *echo.Echo, client *ent.Client) {
	// SEO
	app.GET("/robots.txt", robotsHandler)
	app.GET("/sitemap.xml", sitemapHandler)

	// Routes
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World 👋!")
	})

	app.POST("/users", userCreateHandler(client))
}
