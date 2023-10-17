package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	slogfiber "github.com/samber/slog-fiber"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func server(cmd *cobra.Command, args []string) error {
	app := fiber.New()

	// Default middleware
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(requestid.New())
	app.Use(slogfiber.New(slog.Default())) // Customize logger here

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	eg, ctx := errgroup.WithContext(cmd.Context())
	eg.Go(func() error {
		slog.Info("Running server", "port", "3000")
		return app.Listen(":3000")
	})
	eg.Go(func() error {
		<-ctx.Done()
		slog.Info("Shutting down server")
		return app.Shutdown()
	})

	return eg.Wait()
}
