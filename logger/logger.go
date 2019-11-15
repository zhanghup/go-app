package logger

import (
	"io/ioutil"
	"log"
	"os"
)

type iLogger struct {
	level int
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

type LoggerConfig struct {
	Level int
}

var llog *iLogger

func InitLogger(cfg LoggerConfig) {
	if llog != nil {
		return
	}
	llog = new(iLogger)
	llog.Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	llog.Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	llog.Warn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	//llog.Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
