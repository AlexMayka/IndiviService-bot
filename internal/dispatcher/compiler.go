package dispatcher

// Compiler is a structure that compiles router definitions into an internal map
// of state entries, enabling fast dispatching of handlers by FSM state.
type Compiler struct {
	// States maps FSM state names to compiled handler entries (commands, callbacks, etc.)
	States map[string]*StateEntry
}

// NewCompiler creates and returns a new Compiler instance with initialized state map.
func NewCompiler() *Compiler {
	return &Compiler{
		States: make(map[string]*StateEntry),
	}
}

// GetStates returns the internal map of all compiled FSM states and their handlers.
func (c *Compiler) GetStates() map[string]*StateEntry {
	return c.States
}

// GetState returns a specific StateEntry by name (FSM state).
// The second return value indicates whether the state exists.
func (c *Compiler) GetState(name string) (*StateEntry, bool) {
	entry, ok := c.States[name]
	return entry, ok
}
