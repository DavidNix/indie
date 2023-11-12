package server

import (
	"net/http"

	"github.com/DavidNix/indie/ent"
	"github.com/labstack/echo/v4"
)

func addRoutes(app *echo.Echo, client *ent.Client) {
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World 👋!")
	})
	app.GET("/users", userIndexHandler(client))
	app.POST("/users", userCreateHandler(client))
}
