package server

import (
	"log/slog"
	"time"

	"github.com/DavidNix/indie/ent"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

// NewApp initializes a new server application with predefined middleware and routes.
// It sets read, write and idle timeouts for the server and adds middleware for recovery,
// compression, logging, security headers and CSRF protection.
// Finally, it initializes the routes for the application.
func NewApp(client *ent.Client) *echo.Echo {
	app := echo.New()
	app.Server.ReadTimeout = 60 * time.Second
	app.Server.WriteTimeout = 60 * time.Second
	app.Server.IdleTimeout = 120 * time.Second
	app.IPExtractor = echo.ExtractIPFromXFFHeader() // TODO: Change based on your LB or reverse proxy, https://echo.labstack.com/docs/ip-address.

	app.Use(
		middleware.RequestID(),
		slogecho.New(slog.Default()),
		middleware.Recover(),
		middleware.RemoveTrailingSlash(),
		middleware.CORS(),
		middleware.Gzip(),
		middleware.Decompress(),
		middleware.CSRF(),
		middleware.Secure(),
		middleware.BodyLimit("1M"),
	)

	addRoutes(app, client)

	return app
}

func csrfToken(c echo.Context) string {
	return c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
}

func isHTMX(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}
