package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/DavidNix/indie/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	// Demonstrates config with Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}

	root := cobra.Command{
		RunE: server.server,
	}

	// Run root command
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	err := root.ExecuteContext(ctx)
	cancel()
	if err != nil {
		os.Exit(1)
	}
}
