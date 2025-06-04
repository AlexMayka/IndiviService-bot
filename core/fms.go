package core

// FMS represents a finite state machine for user-specific interaction flows.
// It allows storing and retrieving states and associated key-value parameters per user.
type FMS interface {
	// Get returns the current state of the given user.
	Get(userID int) string

	// Set assigns a new state to the given user.
	Set(userID int, state string)

	// GetParam retrieves a stored parameter for the given user.
	GetParam(userID int, key string) string

	// SetParam sets a parameter value for the given user in the specified state.
	SetParam(userID int, state string, param string)
}
