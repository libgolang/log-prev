# Go Logger

[![GoDoc](https://godoc.org/github.com/libgolang/log?status.svg)](https://godoc.org/github.com/libgolang/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/libgolang/log)](https://goreportcard.com/report/github.com/libgolang/log)
[![License](https://img.shields.io/github/license/libgolang/log.svg)](LICENSE)
[![Release](https://img.shields.io/github/release/libgolang/log.svg)](https://github.com/libgolang/log/releases)





## Download

    go get -u github.com/libgolang/log


## Simple Usage

    package main
    
    import (
    	"github.com/libgolang/log"
    )
    
    func main() {
	
	// debug level
	log.GetDefaultWriter().SetLevel(log.WARN)

	// sets trace
	log.SetTrace(true)

    	log.Debug("This is a debugging statement ... won't show")
    	log.Info("This is a debugging statement  ... won't show")
    	log.Warn("This is a debugging statement  ... will show")
    	log.Error("This is a debugging statement ... will show")
    }

## Configuration Environment Variables

    // Path to configuration file
    export LOG_CONFIG=/path/to/config.properties

    // Override Global Debug Level
    export LOG_LEVEL=DEBUG

### Exampe config.properties

    #
    # Global Debug Level
    # Default: WARN
    log.level=DEBUG

    #
    # Enable Trace. This might be a slow operation
    # Default: false
    log.trace=true

    #
    # Define a log writer. By default there is one writer to stdout
    #
    log.writer.logger0.type=stdout
    
    #
    # Define a log writer to a file
    #
    log.writer.logger1.type=file
    log.writer.logger1.name=one
    log.writer.logger1.level=INFO
    log.writer.logger1.maxfiles=10
    log.writer.logger1.maxSize=1073741824
    log.writer.logger1.dir=./log

