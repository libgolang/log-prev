package logger

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

// WriterChannel implements Writer interface
type WriterChannel struct {
	//writer LoggerWriter
	chn    chan *writerChannelMsg
	writer Writer
}

type writerChannelMsg struct {
	name   string
	mLevel Level
	format string
	args   []interface{}
}

// WriteLog implementation of Writer interface
func (wc *WriterChannel) WriteLog(name string, mLevel Level, format string, args []interface{}) {
	wc.chn <- &writerChannelMsg{name, mLevel, format, args}
}

// NewWriterChannel instantiates a WriterChannel
func NewWriterChannel(writer Writer) Writer {
	wc := &WriterChannel{make(chan *writerChannelMsg, 256), writer}
	go loggerChannelReceiver(wc)
	return wc
}

func loggerChannelReceiver(wc *WriterChannel) {
	var m *writerChannelMsg
	var ok bool
	for {
		m, ok = <-wc.chn
		if !ok {
			break
		}
		wc.writer.WriteLog(m.name, m.mLevel, m.format, m.args)
	}
}
