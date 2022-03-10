package logger

import (
	"log"
	"os"
)

var (
	InfoLogger    = log.New(os.Stderr, "[INFO] ", log.Flags()|log.Lshortfile)
	ErrorLogger   = log.New(os.Stderr, "[ERROR] ", log.Flags())
	FailureLogger = log.New(os.Stderr, "[FAILURE] ", log.Flags())
	SuccessLogger = log.New(os.Stderr, "[SUCCESS] ", log.Flags())
)
