package server

import (
	"github.com/DavidNix/indie/ent"
	"github.com/DavidNix/indie/server/views"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

func userIndexHandler(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := client.User.Query().All(c.Request().Context())
		if err != nil {
			return err
		}
		names := lo.Map(users, func(u *ent.User, _ int) string { return u.Name })
		return views.Render(c, views.ListUsers(names, csrfToken(c)))
	}
}

func userCreateHandler(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("HX-Trigger", "newUser")
		return client.User.Create().SetName(c.FormValue("name")).Exec(c.Request().Context())
	}
}
