package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const logLevelProduction = 5
const logLevelInfo = 4
const logLevelDebug = 3
const logLevelError = 2
const logLevelFatal = 1
const logLevelPanic = 0

var currentLogLevel = logLevelProduction
var shouldPrintWriteLogsInsideProgram = false

type Logger struct {
	logDir         string
	maxLogsPerFile int
	maxLogFiles    int
	currentFile    *os.File
	logCount       int
}

func NewLogger(logDir string, maxLogsPerFile, maxLogFiles int, deletePreviousLogFiles bool) (*Logger, error) {
	if deletePreviousLogFiles {
		// Delete all existing log files in the directory
		logFiles, err := filepath.Glob(filepath.Join(logDir, "log_*.txt"))
		if err != nil {
			return nil, fmt.Errorf("failed to list log files: %w", err)
		}
		for _, file := range logFiles {
			if err := os.Remove(file); err != nil {
				return nil, fmt.Errorf("failed to delete log file %s: %w", file, err)
			}
		}
	}

	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}
	return &Logger{
		logDir:         logDir,
		maxLogsPerFile: maxLogsPerFile,
		maxLogFiles:    maxLogFiles,
	}, nil
}

func (l *Logger) rotateLogs() error {
	files, err := filepath.Glob(filepath.Join(l.logDir, "log_*.txt"))
	if err != nil {
		return fmt.Errorf("failed to list log files: %w", err)
	}

	// Sort files by creation time (using file names)
	if len(files) >= l.maxLogFiles {
		// Remove the oldest log file
		err := os.Remove(files[0])
		if err != nil {
			return fmt.Errorf("failed to remove oldest log file: %w", err)
		}
	}

	// Create a new log file
	newFileName := filepath.Join(l.logDir, fmt.Sprintf("log_%d.txt", time.Now().Unix()))
	file, err := os.Create(newFileName)
	if err != nil {
		return fmt.Errorf("failed to create new log file: %w", err)
	}

	if l.currentFile != nil {
		l.currentFile.Close()
	}
	l.currentFile = file
	l.logCount = 0
	return nil
}

func (l *Logger) PrintlnAndWriteLog(level int, v ...any) {
	l.PrintlnLog(level, v...)
	l.WriteLog(level, v...)
}

func (l *Logger) PrintlnLog(level int, v ...any) {
	if level >= currentLogLevel {
		log.Println(v...)
	}
}

// func (l *Logger) PrintLog(level int, v ...any) {
// 	if level >= currentLogLevel {
// 		log.Print(v...)
// 	}
// }

func (l *Logger) PrintFmt(level int, v ...any) {
	if level >= currentLogLevel {
		fmt.Print(v...)
	}
}

func (l *Logger) PrintlnFmt(level int, v ...any) {
	if level >= currentLogLevel {
		fmt.Println(v...)
	}
}

func (l *Logger) WriteLog(level int, v ...any) error {
	if level >= currentLogLevel {
		if l.currentFile == nil || l.logCount >= l.maxLogsPerFile {
			if err := l.rotateLogs(); err != nil {
				return err
			}
		}

		logEntry := fmt.Sprintf("%s - %s\n", time.Now().Format("2006-01-02T15:04:05.000000000Z07:00"), v)
		_, err := l.currentFile.WriteString(logEntry)
		if err != nil {
			return fmt.Errorf("failed to write to log file: %w", err)
		}
		l.logCount++

		if shouldPrintWriteLogsInsideProgram {
			log.Println(logEntry)
		}
	}

	return nil
}

// func logPrintln(level int, v ...any) {
// 	if level >= currentLogLevel {
// 		log.Println(level, v)
// 	}
// }
