package hermes

// Default is the theme by default
type Default struct{}

// Name returns the name of the default theme
func (dt *Default) Name() string {
	_ = "STUB: not implemented"

	// HTMLTemplate returns a Golang template that will generate an HTML email.
	return ""
}

func (dt *Default) HTMLTemplate() string { _ = "STUB: not implemented"; return "" }

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *Default) PlainTextTemplate() string { _ = "STUB: not implemented"; return "" }
