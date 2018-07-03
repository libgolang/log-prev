package main

import (
	"os"
	"time"

	"github.com/libgolang/log"
)

func main() {
	/*
		w := &log.WriterStdout{}
		w.SetLevel(log.DEBUG)
		log.SetWriters([]log.Writer{w})
	*/
	log.SetTrace(true)
	os.Setenv("LOG_CONFIG", "./config.properties")

	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	time.Sleep(time.Second * 2)
}
