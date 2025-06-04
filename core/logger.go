package core

// Logger defines a unified logging interface used across the SDK.
// It supports structured logging at different severity levels.
type Logger interface {
	// Debug logs a debug-level message.
	Debug(msg string, args ...interface{})

	// Info logs an informational message.
	Info(msg string, args ...interface{})

	// Warn logs a warning message.
	Warn(msg string, args ...interface{})

	// Error logs an error message.
	Error(msg string, args ...interface{})
}
