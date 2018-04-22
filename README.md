# Go Logger




## Download

    go get -u github.com/rhamerica/go-log


## Simple Usage

    package main
    
    import (
    	"github.com/rhamerica/go-log"
    )
    
    func main() {
	
	log.SetDefaultLevel(log.WARN)
	
    	log.Debug("This is a debugging statement ... won't show")
    	log.Info("This is a debugging statement  ... won't show")
    	log.Warn("This is a debugging statement  ... will show")
    	log.Error("This is a debugging statement ... will show")
    }


## Configuration


    package main
    
    import (
    	"github.com/rhamerica/go-log"
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
        "github.com/rhamerica/go-log"
    ) 
     
    func main() {
        l := log.New("main")

        l.Debug("Debug Message")
        l.Info("Info Message")
        l.Warn("Warn Message")
        l.Error("Error Message" )
        l.Panic("Panic Message") // calls panic()
    }



[![Go Report Card](https://goreportcard.com/badge/github.com/rhamerica/go-logger)](https://goreportcard.com/report/github.com/rhamerica/go-logger)
