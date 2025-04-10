package server

import (
	"fmt"
	"html"
	"net/http"

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
		return views.Render(c, views.UserIndex(names, csrfToken(c)))
	}
}

func userCreateHandler(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := client.User.Create().SetName(c.FormValue("name")).Save(c.Request().Context())
		if err != nil {
			return err
		}
		// Beware of XSS.
		escaped := html.EscapeString(user.Name)
		return c.HTML(http.StatusOK, fmt.Sprintf(`<li>%s</li>`, escaped))
	}
}
