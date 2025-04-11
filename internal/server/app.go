package server

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	Environment string
}

func (app *App) Build() {
	mux := echo.New()
	mux.Server.ReadTimeout = 60 * time.Second
	mux.Server.WriteTimeout = 60 * time.Second
	mux.Server.IdleTimeout = 120 * time.Second
	mux.IPExtractor = echo.ExtractIPFromXFFHeader() // TODO: Change based on your LB or reverse proxy, https://echo.labstack.com/docs/ip-address.

	app.registerRoutes(mux)
}

func csrfToken(c echo.Context) string {
	return c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
}
