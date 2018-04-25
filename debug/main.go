package main

import (
	"github.com/rhamerica/go-log"
)

func main() {
	log.SetDefaultLevel(log.WARN)

	log1 := log.New("myLogger")
	log2 := log.New("OtherLogger")

	log.SetLoggerLevels(map[string]log.Level{"myLogger": log.DEBUG})

	log1.Warn("This is a warning statement ... will show")
	log1.Debug("This is a debugging statement ... will show")

	log2.Warn("This is a warning statement ... will show")
	log2.Debug("This is a debugging statement ... won't show")
}
