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
   defaultLogger = NewLogger(LoggerConfig{UseJSON: false})
}

type LoggerConfig struct {
   UseJSON bool
}

func NewLogger(config LoggerConfig) logr.Logger {
   if config.UseJSON {
   	zl := zerolog.New(os.Stderr).With().Timestamp().Logger()
   	return zerologr.New(&zl)
   }
   zl := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05"}).With().Timestamp().Logger()
   return zerologr.New(&zl)
}

func NewContext(ctx context.Context) context.Context {
   return logr.NewContext(ctx, defaultLogger)
}

func FromContext(ctx context.Context) logr.Logger {
   return logr.FromContextOrDiscard(ctx)
}

