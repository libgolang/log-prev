package log

import (
	"fmt"
	"runtime"
	"time"
)

// WriterStdout writes to the standard output
type WriterStdout struct {
}

// WriteLog implementation of logger.Writer
func (*WriterStdout) WriteLog(
	name string,
	mLevel Level,
	format string,
	args []interface{},
) {
	var preFormat string
	if IsTraceEnabled() {
		_, file, line, _ := runtime.Caller(4)
		preFormat = fmt.Sprintf("%s %s [%s] %s:%d %s\n", time.Now().Format(time.RFC3339), mLevel, name, file, line, format)
	} else {
		preFormat = fmt.Sprintf("%s %s [%s] %s\n", time.Now().Format(time.RFC3339), mLevel, name, format)
	}

	fmt.Printf(preFormat, args...)
}

/*
type writerChannelMsg struct {
	name   string
	mLevel Level
	format string
	args   []interface{}
}
*/
