package logging

import (
	"context"
	"log/slog"
	"os"
)

type key int

var loggerKey key

func NewDefaultLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))

}

func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok {
		return NewDefaultLogger()
	}
	return logger
}
