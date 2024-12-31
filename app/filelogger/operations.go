package filelogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var currentLogFileName = GenerateTodayLogFileName()

var ErrorLogFileExists = fmt.Errorf("Log file already exists")

func RotateLog() error {
	newLogFileName := GenerateTodayLogFileName()

	// Close current log file

	// Rename current log file
	_, err := os.Lstat(newLogFileName)
	if os.IsNotExist(err) {
		os.Create(newLogFileName)
		currentLogFileName = newLogFileName
	} else {
		return ErrorLogFileExists
	}
	return nil
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

func GenerateTodayLogFileName() string {
	return fmt.Sprintf("log_%s.log", time.Now().Format("20060102"))
}
