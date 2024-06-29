package main

import (
	"context"

	"github.com/gkwa/farsmell/internal/logger"
)

func main() {
	ctx := logger.NewContext(context.Background())
	doSomething(ctx)

	jsonLogger := logger.NewLogger(logger.LoggerConfig{UseJSON: true})
	ctx = logger.WithLogger(ctx, jsonLogger)
	doSomethingWithJSON(ctx)

	doSomethingElse(ctx)
}

func doSomething(ctx context.Context) {
	log := logger.FromContext(ctx)
	log.Info("Doing something", "step", 1)
}

func doSomethingWithJSON(ctx context.Context) {
	log := logger.FromContext(ctx)
	log.Info("Doing something with JSON", "step", 2)
}

func doSomethingElse(ctx context.Context) {
	defaultCtx := logger.NewContext(context.Background())
	log := logger.FromContext(defaultCtx)
	log.Info("Doing something else", "step", 3)
}
