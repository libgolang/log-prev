package log

import (
	"testing"
)

func TestLoggerLevels(t *testing.T) {

	checkLoggerTest(DEBUG, "test-logger1", "test-msg1", ERROR, t, true)
	checkLoggerTest(DEBUG, "test-logger1", "test-msg1", WARN, t, true)
	checkLoggerTest(DEBUG, "test-logger1", "test-msg1", INFO, t, true)
	checkLoggerTest(DEBUG, "test-logger1", "test-msg1", DEBUG, t, true)

	checkLoggerTest(INFO, "test-logger2", "test-msg2", ERROR, t, true)
	checkLoggerTest(INFO, "test-logger2", "test-msg2", WARN, t, true)
	checkLoggerTest(INFO, "test-logger2", "test-msg2", INFO, t, true)
	checkLoggerTest(INFO, "test-logger2", "test-msg2", DEBUG, t, false)

	checkLoggerTest(WARN, "test-logger3", "test-msg3", ERROR, t, true)
	checkLoggerTest(WARN, "test-logger3", "test-msg3", WARN, t, true)
	checkLoggerTest(WARN, "test-logger3", "test-msg3", INFO, t, false)
	checkLoggerTest(WARN, "test-logger3", "test-msg3", DEBUG, t, false)

	checkLoggerTest(ERROR, "test-logger4", "test-msg4", ERROR, t, true)
	checkLoggerTest(ERROR, "test-logger4", "test-msg4", WARN, t, false)
	checkLoggerTest(ERROR, "test-logger4", "test-msg4", INFO, t, false)
	checkLoggerTest(ERROR, "test-logger4", "test-msg4", DEBUG, t, false)
}

func checkLoggerTest(
	loggerLevel Level,
	givenLoggerName,
	givenLoggerMsg string,
	givenLoggerLevel Level,
	t *testing.T,
	called bool,
) {

	tw := &testWriter{}
	log := getTestLogger(tw, givenLoggerName, loggerLevel)

	if givenLoggerLevel == ERROR {
		log.Error(givenLoggerMsg)
	}
	if givenLoggerLevel == WARN {
		log.Warn(givenLoggerMsg)
	}
	if givenLoggerLevel == INFO {
		log.Info(givenLoggerMsg)
	}
	if givenLoggerLevel == DEBUG {
		log.Debug(givenLoggerMsg)
	}

	if called {
		if givenLoggerName != tw.loggerName {
			t.Errorf("expected logger name: %s, but got: %s", givenLoggerName, tw.loggerName)
		}
		if givenLoggerLevel != tw.loggerLevel {
			t.Errorf("expected logger level: %d, but got: %d", givenLoggerLevel, tw.loggerLevel)
		}
		if givenLoggerMsg != tw.loggerMsg {
			t.Errorf("expected logger name: %s, but got: %s", givenLoggerMsg, tw.loggerMsg)
		}
	} else {
		if "" != tw.loggerName {
			t.Errorf("expected logger name: %s, but got: %s", givenLoggerName, tw.loggerName)
		}
		if 0 != tw.loggerLevel {
			t.Errorf("expected logger level: %d, but got: %d", givenLoggerLevel, tw.loggerLevel)
		}
		if "" != tw.loggerMsg {
			t.Errorf("expected logger name: %s, but got: %s", givenLoggerMsg, tw.loggerMsg)
		}
	}
}

func getTestLogger(tw *testWriter, name string, level Level) Logger {
	log := New(name)
	log.SetLevel(level)
	globalWriters = []Writer{tw}
	return log
}

type testWriter struct {
	//callback    func(name string, mlvl Level, format string, args []interface{})
	loggerName  string
	loggerLevel Level
	loggerMsg   string
	loggerArgs  []interface{}
}

func (tw *testWriter) WriteLog(name string, mlvl Level, format string, args []interface{}) {
	//tw.callback(name, mlvl, format, args)
	tw.loggerName = name
	tw.loggerLevel = mlvl
	tw.loggerMsg = format
	tw.loggerArgs = args
}
