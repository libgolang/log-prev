package main

import (
	"github.com/rhamerica/go-logger"
)

func main() {

	log1 := logger.NewLogger("myLogger")
	log2 := logger.NewLogger("OtherLogger")

	logger.Configuration(logger.Config{
		DefaultLevel: logger.WARN,
		Levels:       map[string]logger.Level{"myLogger": logger.DEBUG},
	})

	log1.Warn("This is a warning statement ... will show")
	log1.Debug("This is a debugging statement ... will show")

	log2.Warn("This is a warning statement ... will show")
	log2.Debug("This is a debugging statement ... won't show")
}
