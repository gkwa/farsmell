package main

import (
	"context"

	"github.com/gkwa/farsmell/internal/logger"
)

func main() {
	ctx := logger.NewContext(context.Background())
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	log := logger.FromContext(ctx)
	log.Info("Doing something", "step", 1)
	doSomethingElse(ctx)
}

func doSomethingElse(ctx context.Context) {
	log := logger.FromContext(ctx)
	log.Info("Doing something else", "step", 2)
}
