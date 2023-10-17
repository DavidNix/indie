package main

import (
	"log/slog"

	"github.com/DavidNix/indie/server"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func runServer(cmd *cobra.Command, args []string) error {
	app := server.New()

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
