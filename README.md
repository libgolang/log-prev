# Go Logger

[![GoDoc](https://godoc.org/github.com/libgolang/log?status.svg)](https://godoc.org/github.com/libgolang/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/libgolang/log)](https://goreportcard.com/report/github.com/libgolang/log)



## Download

    go get -u github.com/libgolang/go-log


## Simple Usage

    package main
    
    import (
    	"github.com/libgolang/log"
    )
    
    func main() {
	
	log.SetDefaultLevel(log.WARN)
	
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
    # Set Debug Level for a logger named `xyzlogger`
    #
    log.level.xyzlogger=WARN

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
    

## Configuration


    package main
    
    import (
    	"github.com/libgolang/log"
    )
    
    func main() {
    
    	log1 := log.New("myLogger")
    	log2 := log.New("OtherLogger")
    
    	log.SetLoggerLevels(map[string]Level{"myLogger": log.DEBUG})
    
    	log1.Warn("This is a warning statement ... will show")
    	log1.Debug("This is a debugging statement ... will show")
    
    	log2.Warn("This is a warning statement ... will show")
    	log2.Debug("This is a debugging statement ... won't show")
    }



## Example

    import(
        "github.com/libgolang/log"
    ) 
     
    func main() {
        l := log.New("main")

        l.Debug("Debug Message")
        l.Info("Info Message")
        l.Warn("Warn Message")
        l.Error("Error Message" )
        l.Panic("Panic Message") // calls panic()
    }


		   
