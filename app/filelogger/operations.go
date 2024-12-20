package filelogger

import (
	"log"
	"os"
	"time"
)

const (
	currentLogFileName = "log.txt"
)

func RotateLog() {
	currentTime := time.Now()
	newLogFileName := "log_" + currentTime.Format("20060102") + ".log"

	// Close current log file

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

func NewFileLogger() (result *FileLogger, err error) {
	result = &FileLogger{}
	logFile, err := os.OpenFile(currentLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	result.FileLogger = log.New(logFile, "", log.LstdFlags)
	return result, nil
}
