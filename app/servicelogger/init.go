package servicelogger

import (
	"fmt"
	"os"

	"github.com/WeeraW/auto-hoyolab-checkin/app/filelogger"
	"github.com/kardianos/service"
)

var Logger service.Logger
var LogFile *filelogger.FileLogger
var LogToFile bool = false

func Info(v ...interface{}) {
	if LogToFile {
		LogFile.Info(v...)
	} else {
		Logger.Info(v...)
	}
}

func Infof(format string, v ...interface{}) {
	if LogToFile {
		LogFile.Infof(format, v...)
	} else {
		Logger.Infof(format, v...)
	}
}

func Error(v ...interface{}) {
	if LogToFile {
		LogFile.Error(v...)
	} else {
		Logger.Error(v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if LogToFile {
		LogFile.Errorf(format, v...)
	} else {
		Logger.Errorf(format, v...)
	}
}

func Warning(v ...interface{}) {
	if LogToFile {
		LogFile.Warning(v...)
	} else {
		Logger.Warning(v...)
	}
}

func Warningf(format string, v ...interface{}) {
	if LogToFile {
		LogFile.Warningf(format, v...)
	} else {
		Logger.Warningf(format, v...)
	}
}

func Fatal(v ...interface{}) {
	Logger.Error(fmt.Sprintf("[FATAL] %s", fmt.Sprint(v...)))
	os.Exit(1)
}

func Debug(v ...interface{}) {
	Logger.Info(fmt.Sprintf("[DEBUG] %s", fmt.Sprint(v...)))
}
