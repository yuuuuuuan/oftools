package oflog

import (
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	log "github.com/charmbracelet/log"
)

var Print *log.Logger

func Init() {
	Print = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "🍪",
	})
	styles := log.DefaultStyles()
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))
	Print.SetStyles(styles)
}
