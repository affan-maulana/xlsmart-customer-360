package logger

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func Init(logLevel string) {
	prefix := ""
	switch logLevel {
	case "debug":
		prefix = "DEBUG "
	case "warn":
		prefix = "WARN "
	case "error":
		prefix = "ERROR "
	default:
		prefix = "INFO "
	}

	Info = log.New(os.Stdout, prefix, log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime)
}
