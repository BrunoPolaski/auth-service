package logger

import (
	"log"
	"os"
	"strings"
)

var (
	Logger   *log.Logger
	LogLevel string
)

func InitLogger() {
	if os.Getenv("LOG_LEVEL") != "" {
		LogLevel = strings.ToLower(os.Getenv("LOG_LEVEL"))
	} else {
		LogLevel = "info"
	}

	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	if LogLevel == "info" || LogLevel == "" {
		Logger.SetPrefix("INFO: ")
		Logger.Println(message)
	}
}

func Warn(message string) {
	if LogLevel == "warn" || LogLevel == "" {
		Logger.SetPrefix("WARNING: ")
		Logger.Println(message)
	}
}

func Error(message string) {
	if LogLevel == "error" || LogLevel == "" {
		Logger.SetPrefix("ERROR: ")
		Logger.Println(message)
	}
}
