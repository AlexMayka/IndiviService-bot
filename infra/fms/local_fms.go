// Package fms implements a simple in-memory finite state machine (FSM)
// used to manage user session states and parameters.
package fms

import "telegram-sdk/core"

// LocalFMS is a local, in-memory implementation of the core.FMS interface.
// It tracks the state and arbitrary key-value parameters for each user by their chat ID.
type LocalFMS struct {
	states map[int]*State
}

// NewLocalFMS creates a new instance of LocalFMS.
// It initializes an internal map to store state objects per user.
func NewLocalFMS() core.FMS {
	return &LocalFMS{
		states: make(map[int]*State),
	}
}

// Get returns the current FSM state for the given user ID.
// If no state is found, it initializes a new state with an empty string.
func (f *LocalFMS) Get(userId int) string {
	if st, ok := f.states[userId]; ok {
		return st.GetState()
	}

	st := NewState()
	f.states[userId] = st
	return st.GetState()
}

// Set assigns a new state string to the specified user ID.
// If the state does not exist, it creates a new one.
func (f *LocalFMS) Set(userId int, state string) {
	if st, ok := f.states[userId]; ok {
		st.SetState(state)
		return
	}

	st := NewState()
	st.SetState(state)
	f.states[userId] = st
}

// GetParam retrieves a parameter value associated with the user's FSM state.
// Returns an empty string if the key or user state is not found.
func (f *LocalFMS) GetParam(userId int, key string) string {
	if st, ok := f.states[userId]; ok {
		if vl, ok := st.GetValue(key); ok {
			return vl
		}
	}
	return ""
}

// SetParam assigns a parameter value to the user's FSM state.
// If the user state is not initialized, the value is ignored.
func (f *LocalFMS) SetParam(userId int, key string, value string) {
	if st, ok := f.states[userId]; ok {
		st.SetValue(key, value)
	}
}
