package logger

import (
    "context"
    "os"
    "github.com/go-logr/logr"
    "github.com/go-logr/zerologr"
    "github.com/rs/zerolog"
)

var defaultLogger logr.Logger

func init() {
    defaultLogger = NewLogger(LoggerConfig{UseJSON: false, LogLevel: 0})
}

type LoggerConfig struct {
    UseJSON  bool
    LogLevel int
}

func NewLogger(config LoggerConfig) logr.Logger {
    var zl zerolog.Logger
    if config.UseJSON {
        zl = zerolog.New(os.Stderr).With().Timestamp().Logger()
    } else {
        zl = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05"}).With().Timestamp().Logger()
    }
    zl = zl.Level(zerolog.Level(config.LogLevel))
    return zerologr.New(&zl)
}

func NewContext(ctx context.Context) context.Context {
    return logr.NewContext(ctx, defaultLogger)
}

func FromContext(ctx context.Context) logr.Logger {
    return logr.FromContextOrDiscard(ctx)
}

func WithLogger(ctx context.Context, logger logr.Logger) context.Context {
    return logr.NewContext(ctx, logger)
}

