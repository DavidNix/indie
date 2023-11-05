package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/DavidNix/indie/ent"
	"github.com/DavidNix/indie/server"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
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

	// TODO: Remove for production, for example demo only.
	if err = ent.Seed(cmd.Context(), client); err != nil {
		return fmt.Errorf("failed seeding database: %w", err)
	}

	app := server.NewApp(client)

	go func() {
		<-cmd.Context().Done()
		slog.Info("Shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = app.Shutdown(ctx)
	}()

	err = app.Start(":3000")
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
