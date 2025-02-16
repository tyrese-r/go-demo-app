package logger

import (
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func InitLogger(logFilePath string) {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	Info = log.New(multiWriter, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(multiWriter, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(multiWriter, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}
