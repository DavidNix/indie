package server

import (
	"errors"
	"log/slog"
	"time"

	"github.com/DavidNix/indie/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

// NewApp initializes a new fiber application with predefined middleware and routes.
// It sets read, write and idle timeouts for the server and adds middleware for recovery,
// compression, logging, security headers and CSRF protection.
// Finally, it initializes the routes for the application.
func NewApp(client *ent.Client) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	})

	// Middleware stack
	app.Use(
		recover.New(recover.Config{
			EnableStackTrace: true,
		}),
		compress.New(),
		slogfiber.New(slog.Default()),
		helmet.New(),
		csrf.New(),
		favicon.New(),
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
