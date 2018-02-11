package logger

// Config logger configuration
type Config struct {
	// DefaultLevel default logging level
	DefaultLevel Level `json:"default"`
	// Levels map of logger names to Levels
	Levels map[string]Level `json:"levels"`
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
	log, ok := config.loggerMap[name]
	if ok {
		return log
	}

	log = &Logger{name: name, level: config.defaultLevel}
	config.loggerMap[name] = log

	// set writer
	// use a WriterChannel wrapping an StdoutWriter
	log.writer = &WriterStdout{}

	return log
}

// Configuration sets the global config
// for all existing loggers and new loggers
func Configuration(conf Config) {
	config.defaultLevel = conf.DefaultLevel

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
		log.level = v
	}
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
