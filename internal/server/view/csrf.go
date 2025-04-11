package view

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func csrfHTMX(c echo.Context) string {
	return fmt.Sprintf(`{"X-CSRF-Token": %q}`, c.Get(middleware.DefaultCSRFConfig.ContextKey).(string))
}

func csrfInput(c echo.Context) templ.Component {
	token := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	return templ.Raw(fmt.Sprintf(`<input type="hidden" name="_csrf" value="%s">`, token))
}
