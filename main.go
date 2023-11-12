package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	// Run root command
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	err := rootCmd().ExecuteContext(ctx)
	cancel()
	if err != nil {
		os.Exit(1)
	}
}
