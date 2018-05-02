package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
	"time"
)

func TestFile(t *testing.T) {
	dir, _ := ioutil.TempDir("", "test")
	defer func() { _ = os.RemoveAll(dir) }()

	fw := NewFileWriter(dir, "test", Megabyte*1, 2)
	fw.WriteLog("log-name", WARN, "Arg1: %s", []interface{}{"arg1"})

	bytes, err := ioutil.ReadFile(path.Join(dir, "test-0.log"))
	if err != nil {
		t.Errorf("Error reading log file: %s", err)
	} else {
		contents := string(bytes)
		if !strings.Contains(contents, "Arg1: arg1") {
			t.Error("Failed to write to file")
		}
	}
}

func TestRotation(t *testing.T) {
	dir, _ := ioutil.TempDir("", "test")
	defer func() { _ = os.RemoveAll(dir) }()

	fw := NewFileWriter(dir, "test", Kilobyte*1, 5)

	msg := "##################### %d"
	for i := 0; i < 100; i++ {
		fw.WriteLog("log-name", WARN, msg, []interface{}{i})
	}

	for i := len(fw.(*fileWriter).logQueue); i > 0; i = len(fw.(*fileWriter).logQueue) {
		t.Logf("Waiting for %d messages...", i)
		time.Sleep(time.Millisecond * 50)
	}

	time.Sleep(time.Second)
	for i := 0; i <= 5; i++ {
		bytes, err := ioutil.ReadFile(path.Join(dir, fmt.Sprintf("test-%d.log", i)))
		if err != nil {
			t.Errorf("Error reading log file: %s", err)
		} else {
			logName := path.Join(dir, fmt.Sprintf("test-%d.log", i))
			stat, err := os.Stat(logName)
			if err == nil {
				t.Logf("logName: %s, size: %d", logName, stat.Size())
			}

			contents := string(bytes)
			if !strings.Contains(contents, "##################### ") {
				t.Errorf("Failed to write to file %s, %d", logName, int64(Kilobyte))
			}
		}
	}

}
