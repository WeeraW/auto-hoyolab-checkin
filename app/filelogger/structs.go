package filelogger

import (
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
)

type FileLogger struct {
	FileLogger *log.Logger
	service.Logger
}

func (l *FileLogger) Error(v ...interface{}) {
	l.FileLogger.Println(v...)
}

func (l *FileLogger) Errorf(format string, v ...interface{}) {
	l.FileLogger.Printf(format, v...)
}

func (l *FileLogger) Info(v ...interface{}) {
	l.FileLogger.Println(v...)
}

func (l *FileLogger) Infof(format string, v ...interface{}) {
	l.FileLogger.Printf(format, v...)
}

func (l *FileLogger) Warning(v ...interface{}) {
	l.FileLogger.Println(v...)
}

func (l *FileLogger) Warningf(format string, v ...interface{}) {
	l.FileLogger.Printf(format, v...)
}

func (l *FileLogger) Debug(v ...interface{}) {
	l.FileLogger.Println(v...)
}

func (l *FileLogger) Debugf(format string, v ...interface{}) {
	l.FileLogger.Printf(format, v...)
}

func (l *FileLogger) Fatal(v ...interface{}) {
	l.FileLogger.Println(v...)
	os.Exit(1)
}

func (l *FileLogger) Fatalf(format string, v ...interface{}) {
	l.FileLogger.Printf(format, v...)
	os.Exit(1)
}

func (l *FileLogger) Close() error {
	return l.FileLogger.Writer().(*os.File).Close()
}

func (l *FileLogger) RotateLog() {
	currentTime := time.Now()
	newLogFileName := "log_" + currentTime.Format("20060102") + ".log"

	// Close current log file
	l.Close()

	// Rename current log file
	err := os.Rename(currentLogFileName, newLogFileName)
	if err != nil {
		// handle error
	}

	// Create new log file
	_, err = os.Create(currentLogFileName)
	if err != nil {
		// handle error
	}
}
