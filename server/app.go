package server

import (
	"errors"
	"log/slog"
	"time"

	"github.com/DavidNix/indie/ent"
	"github.com/gofiber/fiber/v2"
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

	routes(app, client)

	return app
}

func errorHandler(c *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	slog.Error(err.Error(), "status", code)

	if code >= 500 {
		// Obfuscate internal server errors.
		return c.Status(code).SendString(fiber.NewError(code).Error())
	}

	c.Status(code).SendString(e.Error())

	return nil
}
