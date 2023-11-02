package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/DavidNix/indie/ent"
	"github.com/DavidNix/indie/server"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func runServer(cmd *cobra.Command, args []string) error {
	const driver = "sqlite3"
	client, err := ent.Open(driver, os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("failed top open connection to %s: %w", driver, err)
	}
	defer client.Close()

	// Auto-database migrations
	if err = client.Schema.Create(cmd.Context()); err != nil {
		return fmt.Errorf("failed creating schema resources: %w", err)
	}

	app := server.NewApp(client)
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
