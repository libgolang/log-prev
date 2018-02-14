package log

import (
	"fmt"
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
	fmt.Printf(fmt.Sprintf("%s %s [%s] %s\n", time.Now().Format(time.RFC3339), mLevel, name, format), args...)
}

type writerChannelMsg struct {
	name   string
	mLevel Level
	format string
	args   []interface{}
}
