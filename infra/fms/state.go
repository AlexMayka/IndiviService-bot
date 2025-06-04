package fms

// State represents the internal state of a single user in the FSM.
// It contains a string-based state identifier and an optional key-value map
// for storing user-specific parameters (e.g., collected form data).
type State struct {
	state string
	Value map[string]string
}

// NewState initializes and returns a pointer to a new State instance.
// The initial state is an empty string, and the Value map is empty.
func NewState() *State {
	return &State{
		state: "",
		Value: make(map[string]string),
	}
}

// GetValue retrieves a parameter value associated with the given key.
// Returns the value and a boolean indicating whether the key was found.
func (s *State) GetValue(key string) (string, bool) {
	v, ok := s.Value[key]
	return v, ok
}

// SetValue stores a key-value pair in the state’s parameter map.
func (s *State) SetValue(key string, value string) {
	s.Value[key] = value
}

// GetState returns the current state identifier as a string.
func (s *State) GetState() string {
	return s.state
}

// SetState updates the current state identifier.
func (s *State) SetState(state string) {
	s.state = state
}
