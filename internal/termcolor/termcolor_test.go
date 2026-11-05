package termcolor

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestParseSpec(t *testing.T) {
	if _, ok := Parse("default").(lipgloss.NoColor); !ok {
		t.Error(`"default" should map to NoColor`)
	}
	if _, ok := Parse(" Default ").(lipgloss.NoColor); !ok {
		t.Error("default should be case- and space-insensitive")
	}
	if got := Parse("62"); got != lipgloss.Color("62") {
		t.Errorf("Parse(62) = %v", got)
	}
	if got := Parse("#ff0000"); got != lipgloss.Color("#ff0000") {
		t.Errorf("Parse(hex) = %v", got)
	}

	s := ParseSpec("default faint")
	if _, ok := s.Color.(lipgloss.NoColor); !ok || !s.Faint || s.Bold {
		t.Errorf("default faint = %+v", s)
	}
	s = ParseSpec("4 bold underline")
	if s.Color != lipgloss.Color("4") || !s.Bold || !s.Underline || s.Faint {
		t.Errorf("4 bold underline = %+v", s)
	}
	if s := ParseSpec("reverse"); !s.Reverse {
		t.Errorf("reverse = %+v", s)
	}
	if st := ParseSpec("1 faint").Apply(lipgloss.NewStyle()); !st.GetFaint() {
		t.Error("Apply did not set faint")
	}
}
