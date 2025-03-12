package logger

import (
	"log"
	"os"
)

var (
	logger *log.Logger
)

func InitLogger() {
	logger = log.New(os.Stdout, "", log.Lshortfile)
}

func Info(message string) {
	logger.SetPrefix("INFO: ")
	logger.Println(message)
}

func Warn(message string) {
	logger.SetPrefix("WARNING: ")
	logger.Println(message)
}

func Error(message string) {
	logger.SetPrefix("ERROR: ")
	logger.Println(message)
}
