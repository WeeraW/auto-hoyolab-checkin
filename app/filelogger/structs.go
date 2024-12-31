package filelogger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
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

func (l *FileLogger) SetOutput(w *os.File) {
	l.FileLogger.SetOutput(w)
}

func (l *FileLogger) RotateLog() {
	currentTime := time.Now()
	// Close current log file
	l.Close()
	RotateLog()

	// keep only 7 days of logs
	// delete logs older than 7 days
	files, err := os.ReadDir(".")
	if err != nil {
		// handle error
		beeep.Alert("Error", fmt.Sprintf("Error when reading directory: %s", err.Error()), "")
	}
	// at beginning of the day
	oldDate := currentTime.AddDate(0, 0, -7).Truncate(24 * time.Hour)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == currentLogFileName {
			continue
		}
		if strings.Contains(file.Name(), "log_") {
			fileTime, err := time.Parse("20060102", file.Name()[4:12])
			if err != nil {
				// handle error
				beeep.Alert("Error", fmt.Sprintf("Error when parsing file name: %s", err.Error()), "")
			}
			if fileTime.Before(oldDate) {
				err := os.Remove(file.Name())
				if err != nil {
					// handle error
					beeep.Alert("Error", fmt.Sprintf("Error when deleting file: %s", err.Error()), "")
				}
			}
		}
	}
}

// VerifyLogFileAge checks if the log file is not in the same day and rotates the log file
// func (l *FileLogger) VerifyLogFileAge() (err error) {
// 	// Check if log file is older than 24 hours
// 	today := time.Now().Format("20060102")
// 	fileInfo, err := os.Stat(currentLogFileName)
// 	if err != nil {
// 		return err
// 	}
// 	// check if log file is not in the same day
// 	if fileInfo.ModTime().Format("20060102") != today {
// 		l.RotateLog()
// 	}
// 	return nil
// }
