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
		//Prefix:          "🍪",
	})
	styles := log.DefaultStyles()
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFOF🌟").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#90EE9080")).
		Foreground(lipgloss.Color("#006400FF")).Bold(true)

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR🔥").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#FF0000FF")).
		Foreground(lipgloss.Color("#00FFFF00")).Bold(true)

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("FATAL⚡️").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#000000FF")).
		Foreground(lipgloss.Color("#00FFFF00")).Bold(true)

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("Powered by yuuuuuuan🍪").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#000000FF")).
		Foreground(lipgloss.Color("#00FFFF00")).Bold(true)
	Print.SetStyles(styles)
}
