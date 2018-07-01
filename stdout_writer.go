package log

import (
	"fmt"
	"runtime"
	"time"
)

// WriterStdout writes to the standard output
type WriterStdout struct {
	level Level
}

// WriteLog implementation of logger.Writer
func (w *WriterStdout) WriteLog(
	name string,
	mLevel Level,
	format string,
	args []interface{},
) {
	if w.level < mLevel {
		return
	}

	var preFormatStr string
	var preFormatArgs []interface{}
	if IsTraceEnabled() {
		_, file, line, _ := runtime.Caller(4)
		preFormatStr = "%s %s [%s] %s:%d %s\n"
		preFormatArgs = []interface{}{time.Now().Format(time.RFC3339), mLevel.StringColor(), name, file, line, format}
	} else {
		preFormatStr = "%s %s [%s] %s\n"
		preFormatArgs = []interface{}{time.Now().Format(time.RFC3339), mLevel.StringColor(), name, format}
	}

	preFormat := fmt.Sprintf(preFormatStr, preFormatArgs...)
	fmt.Printf(preFormat, args...)
}

// SetLevel sets the writer level
func (w *WriterStdout) SetLevel(level Level) {
	w.level = level
}
