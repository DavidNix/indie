package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppBuilder struct {
	SiteName    string
	PrimaryHost string
	Environment string
}

func (builder AppBuilder) Build() *App {
	mux := echo.New()
	mux.IPExtractor = echo.ExtractIPFromXFFHeader() // TODO: Change based on your LB or reverse proxy, https://echo.labstack.com/docs/ip-address.

	builder.registerRoutes(mux)

	return &App{mux: mux}
}

type App struct {
	mux *echo.Echo
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.mux.ServeHTTP(w, r)
}
