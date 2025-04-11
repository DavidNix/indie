package main

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

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

	app := server.AppBuilder{}.Build()

	fmt.Print("PORT", cmp.Or(os.Getenv("PORT"), "3000"))
	srv := &http.Server{
		Addr:         ":" + cmp.Or(os.Getenv("PORT"), "3000"),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      app,
	}

	var eg errgroup.Group
	eg.Go(func() error {
		slog.Info("Starting server", "address", srv.Addr, "version", version.Version)
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
