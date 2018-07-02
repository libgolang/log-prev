package log

import (
	"encoding/json"
	"github.com/fatih/color"
	"strings"
)

var (
	red    = color.New(color.FgRed).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
	cyan   = color.New(color.FgCyan).SprintfFunc()
	blue   = color.New(color.FgBlue).SprintfFunc()
	//green  = color.New(color.FgGreen).SprintfFunc()
	//magenta = color.New(color.FgMagenta).SprintfFunc()
	//white   = color.New(color.FgWhite).SprintfFunc()
)

const (
	// DEBUG debug level
	DEBUG Level = 40

	// INFO debug level
	INFO Level = 30

	// WARN warn level
	WARN Level = 20

	// ERROR error level
	ERROR Level = 10

	// OTHER empty value
	OTHER Level = 0
)

// Level Logging Level
type Level int

func (l Level) String() string {
	return LevelToStr(l)
}

// StringColor terminal colored string
func (l Level) StringColor() string {
	var str string

	switch l {
	case DEBUG:
		str = blue("DEBUG")
	case INFO:
		str = cyan("INFO")
	case WARN:
		str = yellow("WARN")
	case ERROR:
		str = red("ERROR")
	default:
		str = "OTHER"
	}

	return str
}

// MarshalJSON json serializer
func (l *Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(LevelToStr(*l))
}

// UnmarshalJSON json unserializer
func (l *Level) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	*l = StrToLevel(str)
	return err
}

// StrToLevel converts a string to a Level type
func StrToLevel(str string) Level {
	str = strings.ToUpper(str)
	var level Level
	switch str {
	case "DEBUG":
		level = DEBUG
	case "INFO":
		level = INFO
	case "WARN":
		level = WARN
	case "ERROR":
		level = ERROR
	default:
		level = OTHER
	}
	return level
}

// LevelToStr Level type to a string
func LevelToStr(level Level) string {
	var str string

	switch level {
	case DEBUG:
		str = "DEBUG"
	case INFO:
		str = "INFO"
	case WARN:
		str = "WARN"
	case ERROR:
		str = "ERROR"
	default:
		str = "OTHER"
	}

	return str
}
