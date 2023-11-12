package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Render renders a component to a string and calls c.HTML() with status code 200.
func Render(c echo.Context, t templ.Component) error {
	var buf strings.Builder
	if err := t.Render(c.Request().Context(), &buf); err != nil {
		return fmt.Errorf("failed to render component: %w", err)
	}
	return c.HTML(http.StatusOK, buf.String())
}
