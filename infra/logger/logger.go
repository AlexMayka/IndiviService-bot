// Package logger provides a wrapper around slog.Logger
// to satisfy the core.Logger interface with standard log levels.
package logger

import (
	"log/slog"
)

// Logger is a thin wrapper around slog.Logger that implements the core.Logger interface.
type Logger struct {
	*slog.Logger
}

// NewLogger creates a new Logger instance using the provided slog.Handler.
// The handler defines how logs will be processed and formatted (e.g., JSON or text).
func NewLogger(h slog.Handler) *Logger {
	return &Logger{slog.New(h)}
}

// Debug logs a message at DEBUG level.
// Accepts a message string and optional key-value arguments.
func (l *Logger) Debug(msg string, v ...interface{}) {
	l.Logger.Debug(msg, v...)
}

// Info logs a message at INFO level.
// Accepts a message string and optional key-value arguments.
func (l *Logger) Info(msg string, v ...interface{}) {
	l.Logger.Info(msg, v...)
}

// Warn logs a message at WARN level.
// Accepts a message string and optional key-value arguments.
func (l *Logger) Warn(msg string, v ...interface{}) {
	l.Logger.Warn(msg, v...)
}

// Error logs a message at ERROR level.
// Accepts a message string and optional key-value arguments.
func (l *Logger) Error(msg string, v ...interface{}) {
	l.Logger.Error(msg, v...)
}
