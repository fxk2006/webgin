package util

import (
	"log"
	"os"
	"fmt"
)

const (
	Debug   = iota
	Info
	Warning
)

type Log struct {
	Level    byte
	FileName string
	Logger   *log.Logger
}

func New(level byte, fileName string) *Log {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0664)
	if err != nil {
		log.Fatalln("open file error")
	}
	Logger := log.New(logFile, "[New] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	var flag string
	switch level {
	case Debug:
		flag = " Debug "
		Logger.SetPrefix("[ Debug ]")
	case Info:
		flag = " Info "
		Logger.SetPrefix("[ Info ]")
	case Warning:
		flag = " Warning "
		Logger.SetPrefix("[ Warning ]")
	}
	Logger.Printf("初始化%s日志成功", flag)
	return &Log{level, fileName, Logger}
}
func (l *Log) Debug(v ...interface{}) {
	if Debug >= l.Level {
		l.Logger.SetPrefix("[Debug] ")
		l.Logger.Output(2, fmt.Sprintln(v...))
	}

}
func (l *Log) Info(v ...interface{}) {
	if Info >= l.Level {
		l.Logger.SetPrefix("[INFO] ")
		l.Logger.Output(2, fmt.Sprintln(v...))
	}
}
func (l *Log) Warning(v ...interface{}) {
	if Warning >= l.Level {
		l.Logger.SetPrefix("[Warning] ")
		l.Logger.Output(2, fmt.Sprintln(v...))
	}
}
