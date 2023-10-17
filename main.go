package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

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
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprint(os.Stdout, "Hello, world!\n")
			return err
		},
	}

	// Run root command
	ctx, cancel := signal.NotifyContext(context.Background())
	err := root.ExecuteContext(ctx)
	cancel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
