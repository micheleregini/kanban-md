package tui

import (
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/x/ansi"

	"github.com/antopolskiy/kanban-md/internal/config"
	"github.com/antopolskiy/kanban-md/internal/task"
)

func TestThickCardBorder(t *testing.T) {
	cfg := config.NewDefault("x")
	b := &Board{cfg: cfg, now: time.Now}
	tk := &task.Task{ID: 1, Title: "T", Priority: "low", Status: "todo", Updated: time.Now()}

	thin := ansi.Strip(b.renderCard(tk, false, 20))
	if !strings.Contains(thin, "╭") || strings.Contains(thin, "┃") {
		t.Errorf("default border should be thin and rounded:\n%s", thin)
	}

	cfg.TUI.ThickCardBorder = true
	applyColors(cfg)
	defer applyColors(nil)
	thick := ansi.Strip(b.renderCard(tk, false, 20))
	if !strings.Contains(thick, "┏") || !strings.Contains(thick, "┃") || !strings.Contains(thick, "━") {
		t.Errorf("thick border should have heavy edges and corners:\n%s", thick)
	}
	if got, want := len(strings.Split(thick, "\n")), len(strings.Split(thin, "\n")); got != want {
		t.Errorf("card height changed: %d vs %d", got, want)
	}
}
