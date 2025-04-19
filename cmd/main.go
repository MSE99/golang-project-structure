package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	mainCmd := createMainCliCommand()

	mainCmd.Run(ctx, os.Args)
}
