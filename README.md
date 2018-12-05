# logger
super simple logger for debugging go projects

## Usage
var Level = NONE
func Debug(v ...interface{})
func Debugf(format string, v ...interface{})

func Error(v ...interface{})
func Errorf(format string, v ...interface{})

func Info(v ...interface{})
func Infof(format string, v ...interface{})

func Print(v ...interface{})
func Printf(format string, v ...interface{})

func Println(v ...interface{})
func Printlnf(format string, v ...interface{})

func Warning(v ...interface{})
func Warningf(format string, v ...interface{})


## Log levels

```go
const (
	DEBUG
	INFO
	WARNING
	ERROR
	NONE
)
```

## Example
```go
package main

import "github.com/davidbeesley/logger"

func main() {
	logger.Level = logger.DEBUG
	logger.Debug("Hello")
	logger.Info("Howdy")
	logger.Warning("Hi!")
	logger.Error("Watch out!")
	logger.Println("Good day")

	logger.Level = logger.WARNING
	logger.Debug("Hello")
	logger.Info("Howdy")
	logger.Warning("Hi!")
	logger.Error("Watch out!")
	logger.Println("Good day")

	logger.Level = logger.NONE
	logger.Debug("Hello")
	logger.Info("Howdy")
	logger.Warning("Hi!")
	logger.Error("Watch out!")
	logger.Println("Good day")

}
```
