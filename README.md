# Go Logger




## Download

    go get -u github.com/rhamerica/go-logger


## Configuration


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



## Example

    import(
        "github.com/rhamerica/go-logger"
    ) 
     
    func main() {
        log := logger.NewLogger("main")

        log.Debug("Debug Message")
        log.Info("Info Message")
        log.Warn("Warn Message")
        log.Error("Error Message" )
        log.Panic("Panic Message") // calls panic()
    }

### Structs


    import(
        "github.com/rhamerica/go-logger"
    ) 
     
    type MyStruct struct {
        log logger.Logger
    }
     
    func (ms *MyStruct) Foo(name string) {
        ms.log.Info("Hello %s", name) // same signature as fmt.Printf()
    }


[![Go Report Card](https://goreportcard.com/badge/github.com/rhamerica/go-logger)](https://goreportcard.com/report/github.com/rhamerica/go-logger)
