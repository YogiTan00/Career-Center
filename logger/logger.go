package logger

import (
	"CareerCenter/utils"
	"log"
	"os"
	"strings"
)

type Logger struct {
	log      *log.Logger
	endPoint string
}

func NewLogger(endPoint string) *Logger {
	return &Logger{
		endPoint: endPoint,
		log:      log.New(os.Stdout, "[Logger] ", log.LstdFlags),
	}
}

func (l *Logger) Info(message string) {
	l.log.Println(utils.Color("blue", "[INFO]"), message)
	l.log.Println(utils.Color("green", "[EndPoint]"), l.endPoint)
}

func (l *Logger) Error(err error) {
	l.log.Println(utils.Color("red", "[ERROR]"), err)
	l.log.Println(utils.Color("green", "[EndPoint]"), l.endPoint)
}

func (l *Logger) InfoWithData(message string, data any) {
	l.log.Println(utils.Color("blue", "[INFO]"), strings.Title(message))
	if l.endPoint != "" {
		l.log.Println(utils.Color("green", "[EndPoint]"), l.endPoint)
	}
	if data != nil {
		l.log.Println(utils.Color("yellow", "[Data]"), data)
	}
}
