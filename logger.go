package logger

import (
	"fmt"
	"sync"
)

const (
	_ = iota
	DEBUG
	INFO
	WARNING
	ERROR
	NONE
)

// Terminal Colors
const (
	kNRM = "\x1B[0m"
	kRED = "\x1B[31m" // Error
	kGRN = "\x1B[32m"
	kYEL = "\x1B[33m" // Warning
	kBLU = "\x1B[34m"
	kMAG = "\x1B[35m" // Info
	kCYN = "\x1B[36m" // Debug
	kWHT = "\x1B[37m"
)

// Level is the amount of information to be printed to the console. Default is none.
var Level = NONE

// mux prevents merged printf.
var mux sync.Mutex

func Debug(format string, v ...interface{}) {

	mux.Lock()
	defer mux.Unlock()

	if Level <= DEBUG {
		fmt.Printf(kCYN+"[DEBUG] "+format+kNRM+"\n", v...)
	}
}

func Info(format string, v ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	if Level <= INFO {
		fmt.Printf(kMAG+"[INFO] "+format+kNRM+"\n", v...)
	}
}

func Warning(format string, v ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	if Level <= WARNING {
		fmt.Printf(kYEL+"[WARNING] "+format+kNRM+"\n", v...)
	}
}

func Error(format string, v ...interface{}) {
	mux.Lock()
	defer mux.Unlock()
	if Level <= ERROR {
		fmt.Printf(kRED+"[ERROR] "+format+kNRM+"\n", v...)
	}
}
