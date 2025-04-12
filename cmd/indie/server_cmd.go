package main

import (
	"cmp"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/DavidNix/indie/asset"
	"github.com/DavidNix/indie/internal/server"
	"github.com/DavidNix/indie/internal/version"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func serverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start the HTTP server",
		RunE:  runServer,
	}
	return cmd
}

func runServer(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	asset.SetCacheKey(version.Version())

	app := server.AppBuilder{
		Environment: cmp.Or(os.Getenv("ENVIRONMENT"), "dev"),
		SiteName:    os.Getenv("SITE_NAME"),
		PrimaryHost: os.Getenv("PRIMARY_HOST"),
	}.Build()

	srv := &http.Server{
		Addr:         ":" + cmp.Or(os.Getenv("SERVER_PORT"), "3000"),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      app,
	}

	var eg errgroup.Group

	eg.Go(func() error {
		slog.Info("Starting server", "address", srv.Addr, slog.String("version", version.Version()))
		return srv.ListenAndServe()
	})

	eg.Go(func() error {
		<-ctx.Done()
		slog.Info("Shutting down server")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		return srv.Shutdown(ctx)
	})

	err := eg.Wait()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
