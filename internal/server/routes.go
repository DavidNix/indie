package server

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

func (app *App) registerRoutes(mux *echo.Echo) {
	csrfMiddleware := middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "header:X-CSRF-Token,header:HX-CSRF-Token,form:_csrf",
		CookieName:     "_csrf",
		CookiePath:     "/",
		CookieDomain:   "", // Explicitly set to empty to use current domain only
		CookieMaxAge:   86400,
		CookieSameSite: http.SameSiteStrictMode,
		CookieSecure:   app.Environment != "dev",
		CookieHTTPOnly: true,
		// Example of how to allow payment processor webhooks
		// Skipper: func(c echo.Context) bool {
		// 	return c.Path() == "/stripe" // Skip CSRF for Stripe webhook
		// },
	})

	mux.Use(
		middleware.RequestID(),
		slogecho.New(slog.Default()),
		middleware.Recover(),
		// Warning revisit. It 429s the public resources like images
		// middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(3)),
		middleware.NonWWWRedirect(),
		middleware.RemoveTrailingSlash(),
		middleware.CORS(),
		middleware.Gzip(),
		middleware.Decompress(),
		csrfMiddleware,
		middleware.Secure(),
		middleware.BodyLimit("1M"),
	)

	// SEO
	mux.GET("/robots.txt", robotsHandler)
	mux.GET("/sitemap.xml", sitemapHandler)

	// Routes
	mux.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World 👋!")
	})
}
