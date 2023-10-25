package server

import (
	"log/slog"
	"time"

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
func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	// Middleware stack
	app.Use(
		recover.New(),
		compress.New(),
		slogfiber.New(slog.Default()),
		helmet.New(),
		csrf.New(),
		favicon.New(),
	)

	routes(app)

	return app
}
