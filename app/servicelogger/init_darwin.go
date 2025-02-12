//go:build darwin
// +build darwin

package servicelogger

import (
	"fmt"
	"log"
	"os"

	"github.com/WeeraW/auto-hoyolab-checkin/app/filelogger"
)

var Logger *log.Logger
var LogFile *filelogger.FileLogger
var LogToFile bool = false

func initPlatform() {
	Logger = log.New(os.Stdout, "", log.LstdFlags)
}

func Info(v ...interface{}) {
	if LogToFile {
		LogFile.Info(v...)
	} else {
		Logger.Println(v...)
	}
}

func Infof(format string, v ...interface{}) {
	if LogToFile {
		LogFile.Infof(format, v...)
	} else {
		Logger.Printf(format, v...)
	}
}

func Error(v ...interface{}) {
	if LogToFile {
		LogFile.Error(v...)
	} else {
		// FIXME: This is a temporary fix for the issue where the logger is not logging the error message.
		Logger.Println(v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if LogToFile {
		LogFile.Errorf(format, v...)
	} else {
		// FIXME: This is a temporary fix for the issue where the logger is not logging the error message.
		Logger.Printf(format, v...)
	}
}

func Warning(v ...interface{}) {
	if LogToFile {
		LogFile.Warning(v...)
	} else {
		// FIXME: This is a temporary fix for the issue where the logger is not logging the error message.
		Logger.Println(v...)
	}
}

func Warningf(format string, v ...interface{}) {
	if LogToFile {
		LogFile.Warningf(format, v...)
	} else {
		// FIXME: This is a temporary fix for the issue where the logger is not logging the error message.
		Logger.Printf(format, v...)
	}
}

func Fatal(v ...interface{}) {
	Logger.Fatalf("[FATAL] %s", fmt.Sprint(v...))
	os.Exit(1)
}

func Debug(v ...interface{}) {
	Logger.Printf("[DEBUG] %s", fmt.Sprint(v...))
}
