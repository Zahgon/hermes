package hermes

// Flat is a theme
type Flat struct{}

// Name returns the name of the flat theme
func (dt *Flat) Name() string {
	_ = "STUB: not implemented"

	// HTMLTemplate returns a Golang template that will generate an HTML email.
	return ""
}

func (dt *Flat) HTMLTemplate() string { _ = "STUB: not implemented"; return "" }

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *Flat) PlainTextTemplate() string { _ = "STUB: not implemented"; return "" }
