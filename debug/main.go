package main

import (
	"time"

	"github.com/libgolang/log"
)

func main() {
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	time.Sleep(time.Second * 2)
}
