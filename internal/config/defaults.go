// Package config handles kanban board configuration.
package config

const (
	// DefaultDir is the default kanban directory name.
	DefaultDir = "kanban"
	// DefaultTasksDir is the default tasks subdirectory name.
	DefaultTasksDir = "tasks"
	// DefaultStatus is the default status for new tasks.
	DefaultStatus = "backlog"
	// DefaultPriority is the default priority for new tasks.
	DefaultPriority = "medium"
	// DefaultClass is the default class of service for new tasks.
	DefaultClass = "standard"
	// DefaultClaimTimeout is the default claim expiration as a duration string.
	DefaultClaimTimeout = "1h"
	// DefaultTitleLines is the default number of title lines in TUI cards.
	DefaultTitleLines = 2
	// DefaultHideEmptyColumns controls whether TUI hides empty status columns.
	DefaultHideEmptyColumns = false

	// ConfigFileName is the name of the config file within the kanban directory.
	ConfigFileName = "config.yml"

	// CurrentVersion is the current config schema version.
	CurrentVersion = 12

	// ArchivedStatus is the reserved status name for soft-deleted tasks.
	ArchivedStatus = "archived"
)

// Default slice values for a new board (slices cannot be const).
var (
	DefaultStatuses = []StatusConfig{
		{Name: "backlog", ShowDuration: boolPtr(false)},
		{Name: "todo"},
		{Name: "in-progress", RequireClaim: true},
		{Name: "review", RequireClaim: true},
		{Name: "done", ShowDuration: boolPtr(false)},
		{Name: ArchivedStatus, ShowDuration: boolPtr(false)},
	}

	DefaultPriorities = []string{
		"low",
		"medium",
		"high",
		"critical",
	}

	// DefaultAgeThresholds defines the default progressive color thresholds
	// for task duration display in the TUI. Tasks are colored based on how
	// long they've been in their current status.
	DefaultAgeThresholds = []AgeThreshold{
		{After: "0s", Color: "242"},   // dim gray (fresh)
		{After: "1h", Color: "34"},    // green
		{After: "24h", Color: "226"},  // yellow
		{After: "72h", Color: "208"},  // orange
		{After: "168h", Color: "196"}, // red (1 week)
	}

	// DefaultTUIColors defines the default TUI colors (ANSI 256 codes or hex
	// values). Keys are the valid names for tui.colors in config.yml.
	DefaultTUIColors = map[string]string{
		"column_header_fg":  "252",
		"column_header_bg":  "236",
		"active_header_fg":  "230",
		"active_header_bg":  "62",
		"drop_target_fg":    "232",
		"drop_target_bg":    "42",
		"card_border":       "240",
		"active_border":     "62",
		"blocked_border":    "196",
		"text":              "252",
		"dim":               "241",
		"error":             "196",
		"priority_critical": "196",
		"priority_high":     "208",
		"priority_medium":   "226",
		"priority_low":      "242",
		"tab_fg":            "250",
		"tab_bg":            "238",
		"claim":             "44",
		// Used by table output (list, show) only.
		"list_header": "244",
		"tag":         "110",
	}

	// DefaultClasses defines the default classes of service.
	DefaultClasses = []ClassConfig{
		{Name: "expedite", WIPLimit: 1, BypassColumnWIP: true},
		{Name: "fixed-date"},
		{Name: "standard"},
		{Name: "intangible"},
	}
)

// boolPtr returns a pointer to the given bool value.
func boolPtr(v bool) *bool { return &v }
