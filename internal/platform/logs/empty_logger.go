package logs

import (
	"context"
	"log/slog"
)

type EmptyLoggerHandler struct{}

func NewEmptyLogger() *slog.Logger {
	return slog.New(NewEmptyLoggerHandler())
}

func NewEmptyLoggerHandler() *EmptyLoggerHandler {
	return &EmptyLoggerHandler{}
}

// Enabled always returns false, since the log entry is ignored
func (EmptyLoggerHandler) Enabled(context.Context, slog.Level) bool {
	return false
}

// Handle just ignores the log entry
func (EmptyLoggerHandler) Handle(context.Context, slog.Record) error {
	panic("implement me")
}

// WithAttrs just returns the same handler, since there are no attributes to save
func (h EmptyLoggerHandler) WithAttrs([]slog.Attr) slog.Handler {
	return h
}

// WithGroup returns the same handler
func (h EmptyLoggerHandler) WithGroup(string) slog.Handler {
	return h
}
