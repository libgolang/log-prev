package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/magiconair/properties"
)

var (
	defaultWriter      = getDefaultWriter()
	globalWriters      = []Writer{defaultWriter}
	globalTraceEnabled = false
)

// Writer writer interface
type Writer interface {
	WriteLog(logLevel Level, format string, args []interface{})
	SetLevel(level Level)
}

// SetTrace when set to true, the log will print file names and line numbers
func SetTrace(trace bool) {
	globalTraceEnabled = trace
}

// IsTraceEnabled whether printing of file names and line numbers is enabled
func IsTraceEnabled() bool {
	return globalTraceEnabled
}

// SetWriters sets the writers for all the loggers
func SetWriters(w []Writer) {
	globalWriters = w
}

// PrintLog sends a log message to the writers.
// name: logger name
// loggerLevel: the level of the logger implementation
// logLevel: the level of the message. If the level of the message is greater than loggerLevel the log will bi discarted
// format: log format.  See fmt.Printf
// a...: arguments.  See fmt.Printf
func PrintLog(logLevel Level, format string, a []interface{}) {
	for _, w := range globalWriters {
		w.WriteLog(logLevel, format, a)
	}
}

// Err log to root logger
func Err(format string, a ...interface{}) error {
	PrintLog(ERROR, format, a)
	return fmt.Errorf(format, a...)
}

// Error log to root logger
func Error(format string, a ...interface{}) {
	PrintLog(ERROR, format, a)
}

// Info log to root logger
func Info(format string, a ...interface{}) {
	PrintLog(INFO, format, a)
}

// Warn log to root logger
func Warn(format string, a ...interface{}) {
	PrintLog(WARN, format, a)
}

// Debug log to root logger
func Debug(format string, a ...interface{}) {
	PrintLog(DEBUG, format, a)
}

// Panic log to root logger
func Panic(format string, a ...interface{}) {
	PrintLog(ERROR, format, a)
	panic(fmt.Errorf(format, a...))
}

// resolves configuration
func getDefaultWriter() Writer {
	return &WriterStdout{WARN}
}

// GetDefaultWriter gets the default writer
func GetDefaultWriter() Writer {
	return defaultWriter
}

// LoadLogProperties loads properties from configuration file in LOG_CONFIG
func LoadLogProperties() {
	cfgFile, ok := os.LookupEnv("LOG_CONFIG")
	if !ok {
		return
	}

	props, err := properties.LoadFile(cfgFile, properties.UTF8)
	if err != nil {
		return
	}

	//
	// Trace
	//
	if props.GetString("log.trace", "false") == "true" {
		SetTrace(true)
	}

	//
	// Writers
	//
	logWriters := make([]Writer, 0)
	processed := make(map[string]bool)
	for k := range props.Map() {
		if strings.HasPrefix(k, "log.writer.") {
			parts := strings.Split(k, ".")
			if len(parts) != 4 {
				continue
			}
			loggerName := parts[2]

			// already set
			if _, ok := processed[loggerName]; ok {
				continue
			}

			var writer Writer
			loggerType := props.GetString(fmt.Sprintf("log.writer.%s.type", loggerName), "stdout")
			loggerLevel := StrToLevel(props.GetString(fmt.Sprintf("log.writer.%s.level", loggerName), "DEBUG"))
			if loggerType == "stdout" {
				writer = &WriterStdout{level: loggerLevel}
			} else if loggerType == "file" {
				size := props.GetInt64(fmt.Sprintf("log.writer.%s.maxSize", loggerName), int64(Gigabyte))
				maxfiles := props.GetInt(fmt.Sprintf("log.writer.%s.maxFiles", loggerName), 10)
				dir := props.GetString(fmt.Sprintf("log.writer.%s.dir", loggerName), "./log")
				name := props.GetString(fmt.Sprintf("log.writer.%s.name", loggerName), loggerName)
				writer = NewFileWriter(dir, name, FileSize(size), maxfiles, loggerLevel)
			} else {
				continue
			}
			processed[loggerName] = true
			logWriters = append(logWriters, writer)
		}
	}

	if len(logWriters) > 0 {
		SetWriters(logWriters)
	}
}

func preformat(
	writerLevel Level,
	messageLevel Level,
	format string,
	do func(preFormat string),
) {
	if writerLevel < messageLevel {
		return
	}

	var preFormatStr string
	var preFormatArgs []interface{}
	if IsTraceEnabled() {
		_, file, line, _ := runtime.Caller(4)
		preFormatStr = "%s %s %s:%d %s\n"
		preFormatArgs = []interface{}{time.Now().Format(time.RFC3339), messageLevel.StringColor(), file, line, format}
	} else {
		preFormatStr = "%s %s %s\n"
		preFormatArgs = []interface{}{time.Now().Format(time.RFC3339), messageLevel.StringColor(), format}
	}

	do(fmt.Sprintf(preFormatStr, preFormatArgs...))
}
