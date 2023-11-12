package http

import (
	"net/http"

	"github.com/DavidNix/indie/ent"
	"github.com/labstack/echo/v4"
)

func userIndexHandler(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := client.User.Query().All(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
}
