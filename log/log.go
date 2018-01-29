package log

import (
	kLog "github.com/zuiwuchang/king-go/log"
	"log"
)

var Trace *log.Logger
var Debug *log.Logger
var Info *log.Logger
var Warn *log.Logger
var Error *log.Logger
var Fault *log.Logger

func init() {
	loggers := kLog.NewDebugLoggers()

	Trace = loggers.Trace
	Debug = loggers.Debug
	Info = loggers.Info
	Warn = loggers.Warn
	Error = loggers.Error
	Fault = loggers.Fault
}
