package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mse99/golang-project-structure/config"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	config.Load()

	mainCmd := createMainCliCommand()

	mainCmd.Run(ctx, os.Args)
}
