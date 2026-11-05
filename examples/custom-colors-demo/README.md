# Custom colors demo

No color here is a fixed value: everything comes from the terminal theme.

- ANSI `0`-`15` are the theme's own palette (red, green, blue, magenta for
  priorities, claim, tags and blocked cards).
- `default` is the terminal's own foreground/background.
- `faint` dims text (ids, tags, dates, low priority, status bar) without
  choosing a gray, so it works on any theme.
- The selected column header and the selected card use the same selection color
  (`active_header_bg` and `active_border`, both `${KBMD_ACCENT:-4}`), so the
  highlight is uniform.
- `reverse` swaps foreground and background; it is not used here because it
  would not match the card selection color.

Yellow, cyan and the gray/white codes (`3`, `6`, `7`, `8`) are avoided for text
because many themes render them too light or too dark; `8` is only used for
card borders, where low contrast is acceptable. The selection color can be
overridden with `KBMD_ACCENT`.

Card borders are heavy with square corners (`thick_card_border: true`). Set it to
`false` for the thin border.

`tui.status_colors` is intentionally not set: it would color each TUI column
header differently. With it unset, headers use one color for the active column
and one for the others, and `list` keeps its default status colors.

```bash
# From the repository root
go run ./cmd/kanban-md --dir examples/custom-colors-demo tui
go run ./cmd/kanban-md --dir examples/custom-colors-demo list

# Change the selection color from the environment
KBMD_ACCENT=5 go run ./cmd/kanban-md --dir examples/custom-colors-demo tui

# Plain output
NO_COLOR=1 go run ./cmd/kanban-md --dir examples/custom-colors-demo list
```

Things to look for: critical priority and `@alice` claim (task 1), tag and
due colors (task 2), blocked border (task 5),
and the narrow tab strip when the terminal is under about 80 columns. Dragging a
card with the mouse shows the drop target header.

Running the CLI on this directory may write `activity.jsonl`; do not commit it.
