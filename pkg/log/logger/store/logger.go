package store

import (
	"context"

	"github.com/setcreed/onexstack/pkg/log"
)

// Logger is a logger that implements the Logger interface.
// It uses the log package to log error messages with additional context.
type Logger struct{}

// NewLogger creates and returns a new instance of Logger.
func NewLogger() *Logger {
	return &Logger{}
}

// Error logs an error message with the provided context using the log package.
func (l *Logger) Error(ctx context.Context, err error, msg string, kvs ...any) {
	log.W(ctx).Errorw(err, msg, kvs...)
}
