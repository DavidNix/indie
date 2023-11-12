package server

import (
	_ "embed"
	"net/http"

	"github.com/DavidNix/indie/ent"
	"github.com/labstack/echo/v4"
)

//go:embed views/robots.txt
var robotstxt string

func addRoutes(app *echo.Echo, client *ent.Client) {
	// SEO
	// TODO: Change the sitemap URL.
	app.GET("/robots.txt", func(c echo.Context) error { return c.String(http.StatusOK, robotstxt) })

	// Routes
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World 👋!")
	})

	app.POST("/users", userCreateHandler(client))
}
