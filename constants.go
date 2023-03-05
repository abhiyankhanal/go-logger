package go_logger

var (
	colors map[LogLevel]string

	// Map from format's placeholders to printf verbs
	phfs map[string]string

	// Contains color strings for stdout
	logNo uint64

	// Default format of log message
	defFmt = "#%[1]d %[2]s %[4]s:%[5]d ▶ %.3[6]s %[7]s"

	// Default format of time
	defTimeFmt = "2006-01-02 15:04:05"
)

// LogLevel type
type LogLevel int

// Log Level
const (
	CriticalLevel LogLevel = iota + 1
	ErrorLevel
	WarningLevel
	NoticeLevel
	InfoLevel
	DebugLevel
)

// Color numbers for stdout
const (
	Black = (iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Constant color maps

var (
	defaultColors = map[LogLevel]string{
		CriticalLevel: ColorString(Magenta),
		ErrorLevel:    ColorString(Red),
		WarningLevel:  ColorString(Yellow),
		NoticeLevel:   ColorString(Green),
		DebugLevel:    ColorString(Cyan),
		InfoLevel:     ColorString(White),
	}
	minimalColors = map[LogLevel]string{
		CriticalLevel: ColorString(Red),
		ErrorLevel:    ColorString(Red),
		WarningLevel:  ColorString(Yellow),
		NoticeLevel:   ColorString(Green),
		DebugLevel:    ColorString(Green),
		InfoLevel:     ColorString(Green),
	}
)
