package log
import(
	log "github.com/charmbracelet/log"
)

logger := log.NewWithOptions(os.Stderr, log.Options{
	ReportCaller:    true,
	ReportTimestamp: true,
	TimeFormat:      time.DateTime,
	//Prefix:          "Baking 🍪 ",
})
