package servicelogger

import (
	"fmt"
	"os"

	"github.com/kardianos/service"
)

var Logger service.Logger

func Info(v ...interface{}) {
	Logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	Logger.Infof(format, v...)
}

func Error(v ...interface{}) {
	Logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	Logger.Errorf(format, v...)
}

func Warning(v ...interface{}) {
	Logger.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	Logger.Warningf(format, v...)
}

func Fatal(v ...interface{}) {
	Logger.Error(v...)
	os.Exit(1)
}

func Debug(v ...interface{}) {
	Logger.Info(fmt.Sprintf("[DEBUG] %s", fmt.Sprint(v...)))
}
