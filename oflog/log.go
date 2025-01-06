package oflog

import (
	"os"
	"time"

	log "github.com/charmbracelet/log"
)

var Print *log.Logger

func Init() {
	Print = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
	})
}
