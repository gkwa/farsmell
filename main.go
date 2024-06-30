package main

import (
	"context"

	"github.com/gkwa/farsmell/internal/logger"
)

func main() {
	// Set log level to Info (0) or Debug (-1) here
	ctx := logger.NewContext(context.Background())
	doSomething(ctx)

	jsonLogger := logger.NewLogger(logger.LoggerConfig{UseJSON: true, LogLevel: 0})
	ctx = logger.WithLogger(ctx, jsonLogger)
	doSomethingWithJSON(ctx)
	doSomethingElse(ctx)
}

func doSomething(ctx context.Context) {
	log := logger.FromContext(ctx)
	log.Info("Doing something", "step", 1)
	log.V(1).Info("Debug info", "details", "some debug details")
}

func doSomethingWithJSON(ctx context.Context) {
	log := logger.FromContext(ctx)
	log.Info("Doing something with JSON", "step", 2)
	log.V(1).Info("Debug info in JSON", "details", "some JSON debug details")
}

func doSomethingElse(ctx context.Context) {
	defaultCtx := logger.NewContext(context.Background())
	log := logger.FromContext(defaultCtx)
	log.Info("Doing something else", "step", 3)
}
