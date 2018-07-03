package log

import (
	"fmt"
)

// WriterStdout writes to the standard output
type WriterStdout struct {
	level Level
}

// NewStdoutWriter constructor
func NewStdoutWriter(level Level) Writer {
	return &WriterStdout{level}
}

// WriteLog implementation of logger.Writer
func (w *WriterStdout) WriteLog(
	messageLevel Level,
	format string,
	args []interface{},
) {
	preformat(w.level, messageLevel, format, func(preFormat string) {
		fmt.Printf(preFormat, args...)
	})
}

// SetLevel sets the writer level
func (w *WriterStdout) SetLevel(level Level) {
	w.level = level
}
