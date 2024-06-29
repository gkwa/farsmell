// main.go
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

// LoggerConfig holds the configuration for the logger
type LoggerConfig struct {
	UseJSON bool
}

// NewLogger creates a new logger based on the provided configuration
func NewLogger(config LoggerConfig) logr.Logger {
	if config.UseJSON {
		zl := zerolog.New(os.Stderr).With().Timestamp().Logger()
		return zerologr.New(&zl)
	}

	// Text console logger
	zl := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05"}).With().Timestamp().Logger()
	return zerologr.New(&zl)
}

func main() {
	// Configure and create the logger
	config := LoggerConfig{
		UseJSON: false, // Set to true for JSON logging, false for text console logging
	}
	Logger = NewLogger(config)

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
