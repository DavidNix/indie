package server

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func New() *fiber.App {
	app := fiber.New()

	// Middleware stack
	app.Use(
		recover.New(),
		compress.New(),
		slogfiber.New(slog.Default()),
		helmet.New(),
		csrf.New(),
	)

	routes(app)

	return app
}
