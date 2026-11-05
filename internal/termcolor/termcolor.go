// Package termcolor converts configured color strings to lipgloss styles.
package termcolor

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Default is the color value that selects the terminal's own foreground or
// background color instead of forcing one, so it adapts to any theme.
const Default = "default"

// Spec is a parsed color value: a color plus optional text attributes.
type Spec struct {
	Color     lipgloss.TerminalColor
	Bold      bool
	Faint     bool
	Italic    bool
	Underline bool
	Reverse   bool
}

// ParseSpec parses a space-separated value such as "62", "#5f5fd7",
// "default faint" or "reverse". The first token that is not an attribute
// (bold, faint, italic, underline, reverse) is the color; "default" or no
// color leaves the terminal's own color in place.
func ParseSpec(v string) Spec {
	s := Spec{Color: lipgloss.NoColor{}}
	for _, tok := range strings.Fields(v) {
		switch strings.ToLower(tok) {
		case "bold":
			s.Bold = true
		case "faint":
			s.Faint = true
		case "italic":
			s.Italic = true
		case "underline":
			s.Underline = true
		case "reverse":
			s.Reverse = true
		case Default:
			s.Color = lipgloss.NoColor{}
		default:
			s.Color = lipgloss.Color(tok)
		}
	}
	return s
}

// Parse returns only the color of v. Use it where attributes cannot apply
// (backgrounds and borders).
func Parse(v string) lipgloss.TerminalColor { return ParseSpec(v).Color }

// Apply sets the spec's foreground color and attributes on st. Attributes that
// are not set in the spec are left as they are in st.
func (s Spec) Apply(st lipgloss.Style) lipgloss.Style {
	st = st.Foreground(s.Color)
	if s.Bold {
		st = st.Bold(true)
	}
	if s.Faint {
		st = st.Faint(true)
	}
	if s.Italic {
		st = st.Italic(true)
	}
	if s.Underline {
		st = st.Underline(true)
	}
	if s.Reverse {
		st = st.Reverse(true)
	}
	return st
}
