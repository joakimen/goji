package format

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	keyColor     = color.New(color.FgCyan)
	summaryColor = color.New(color.FgWhite)
	dateColor    = color.New(color.FgGreen)
)

func StatusColor(status string) string {
	switch strings.ToLower(status) {
	case "done", "resolved", "closed":
		return color.New(color.FgGreen).Sprint(status)
	case "in progress", "started":
		return color.New(color.FgYellow).Sprint(status)
	case "blocked", "impediment":
		return color.New(color.FgRed).Sprint(status)
	default:
		return color.New(color.FgBlue).Sprint(status)
	}
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatItem(key, status, summary string, created time.Time) string {
	return fmt.Sprintf("%s %s %s (%s)",
		keyColor.Sprint(key),
		StatusColor(status),
		summaryColor.Sprint(truncate(summary, 60)),
		dateColor.Sprint(FormatDate(created)))
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}
