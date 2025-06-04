package core

// Regex defines a handler associated with a regular expression pattern.
// It is used in the router to match incoming messages against regex rules.
type Regex struct {
	Pattern string  // Regular expression pattern to match message text
	Method  Handler // Handler to execute if the pattern matches
}

// MsgEquals defines a handler triggered by an exact message match.
// It is used in the router for matching static message content.
type MsgEquals struct {
	Text    string  // Exact message text to match
	Handler Handler // Handler to execute on match
}
