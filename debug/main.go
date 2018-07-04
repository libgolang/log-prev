package main

import (
	"time"

	"github.com/libgolang/log"
)

func main() {
	/*
		w := &log.WriterStdout{}
		w.SetLevel(log.DEBUG)
		log.SetWriters([]log.Writer{w})
	*/
	log.GetDefaultWriter().SetLevel(log.DEBUG)
	log.SetTrace(true)

	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	time.Sleep(time.Second * 2)
}
