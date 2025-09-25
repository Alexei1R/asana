package main

import (
	"asana/internal/application"
	"context"
	"os/signal"
	"syscall"
)

func main() {

	app := application.New()

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	if err := app.Run(ctx); err != nil {
		panic(err)
	}

}
