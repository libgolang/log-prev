# Go Logger




## Download

    go get -u github.com/rhamerica/go-logger

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