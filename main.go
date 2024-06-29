package main

import (
	"context"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

// Logger is a global variable that holds the logger instance
var Logger logr.Logger

func init() {
	// Initialize the logger
	zl := zerolog.New(os.Stderr).With().Timestamp().Logger()
	Logger = zerologr.New(&zl)
}

func main() {
	// Create a context with the logger
	ctx := logr.NewContext(context.Background(), Logger)

	// Call a function that uses the logger
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	// Retrieve the logger from the context
	log := logr.FromContextOrDiscard(ctx)

	log.Info("Doing something", "step", 1)

	// Call another function, passing the context
	doSomethingElse(ctx)
}

func doSomethingElse(ctx context.Context) {
	// Retrieve the logger from the context
	log := logr.FromContextOrDiscard(ctx)

	log.Info("Doing something else", "step", 2)
}
