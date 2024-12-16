package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	logDir := "./logs"
	maxLogsPerFile := 100
	maxLogFiles := 5
	deletePreviousLogFiles := true

	logger, err := NewLogger(
		logDir,
		maxLogsPerFile,
		maxLogFiles,
		deletePreviousLogFiles,
	)
	if err != nil {
		logger.PrintlnAndWriteLog(logLevelFatal, "Failed to initialize logger: ", err)
		return
	}

	if len(os.Args) < 2 {
		logger.PrintlnAndWriteLog(logLevelFatal, "No command provided")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	command := os.Args[1]
	switch command {
	case kCommandCreate:
		if len(os.Args) < 3 {
			logger.PrintlnAndWriteLog(logLevelFatal, " '"+kCommandCreate+"' command requires a command type")
			return
		}
		commandType := os.Args[2]
		switch commandType {
		case kCommandTypeCreateServer:
			logger.PrintFmt(logLevelProduction, "Enter server name: ")
			getInputString(scanner)
			logger.PrintFmt(logLevelProduction, "Enter server display name: ")
		}
		
	}
}
