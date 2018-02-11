package logger

import (
	"strings"
)

// Level Logging Level
type Level int

const (
	// DEBUG debug level
	DEBUG Level = 40

	// INFO debug level
	INFO Level = 30

	// WARN warn level
	WARN Level = 20

	// ERROR error level
	ERROR Level = 10
)

// Config logger configuration
type Config struct {
	DefaultLevel string            `json:"default"`
	Levels       map[string]string `json:"levels"`
	//loggers      map[string]*Logger
}

// Writer writer interface
type Writer interface {
	WriteLog(name string, mlvl Level, format string, args []interface{})
}

// Logger Logging Object
type Logger struct {
	name   string
	level  Level
	writer Writer
}

//
// global configuraction
type configStruct struct {
	defaultLevel Level
	loggerMap    map[string]*Logger
	writer       Writer
}

var config = configStruct{
	defaultLevel: WARN,
	loggerMap:    make(map[string]*Logger),
	writer:       &WriterStdout{},
}

// NewLogger instantiates logger
func NewLogger(name string) *Logger {
	var log *Logger
	log, ok := config.loggerMap[name] //config.loggers[name]
	if ok {
		return log
	}

	log = &Logger{name: name, level: config.defaultLevel}
	config.loggerMap[name] = log

	// set writer
	// use a WriterChannel wrapping an StdoutWriter
	log.writer = NewWriterChannel(&WriterStdout{})

	return log
}

// Configuration sets the global config
// for all existing loggers and new loggers
func Configuration(conf *Config) {
	config.defaultLevel = strToLevel(conf.DefaultLevel)

	// reset existing logger levels
	for _, log := range config.loggerMap {
		log.level = config.defaultLevel
	}

	// Create instances defined in configuration
	for k, v := range conf.Levels {
		log, ok := config.loggerMap[k]
		if !ok {
			log = NewLogger(k)
		}
		log.level = strToLevel(v)
	}
}

func strToLevel(str string) Level {
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
		level = 0
	}
	return level
}

// Error logs at error level
func (log *Logger) Error(format string, a ...interface{}) {
	printLogMessage(log, ERROR, format, a)
}

// Info logs at info level
func (log *Logger) Info(format string, a ...interface{}) {
	printLogMessage(log, INFO, format, a)
}

// Warn logs at wanr level
func (log *Logger) Warn(format string, a ...interface{}) {
	printLogMessage(log, WARN, format, a)
}

// Debug logs at debug level
func (log *Logger) Debug(format string, a ...interface{}) {
	printLogMessage(log, DEBUG, format, a)
}

// Panic error and exit
func (log *Logger) Panic(format string, a ...interface{}) {
	printLogMessage(log, ERROR, format, a)
	panic("panic!")

}

func printLogMessage(log *Logger, methodLevel Level, format string, a []interface{}) {
	if log.level < methodLevel {
		return
	}
	log.writer.WriteLog(log.name, methodLevel, format, a)
}

func (l Level) String() string {
	var lvl string
	if l == ERROR {
		lvl = "ERROR"
	} else if l == WARN {
		lvl = "WARN"
	} else if l == INFO {
		lvl = "INFO"
	} else if l == DEBUG {
		lvl = "DEBUG"
	} else {
		lvl = "OTHER"
	}
	return lvl
}
