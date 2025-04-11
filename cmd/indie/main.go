package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/DavidNix/indie/internal/version"
	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{
		Short:   "My awesome app",
		Version: version.Version,
	}
	root.AddCommand(
		serverCmd(),
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	err := root.ExecuteContext(ctx)
	cancel()
	if err != nil {
		os.Exit(1)
	}
}
