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
	l.log.Println("[INFO]", message)
}

func (l *Logger) Error(message string) {
	l.log.Println("[ERROR]", message)
}

func (l *Logger) General(message string, data any) {
	//If message nil only can handle error
	if message != "" {
		l.log.Println(utils.Color("blue", "[INFO]"), strings.Title(message))
		if l.endPoint != "" {
			l.log.Println(utils.Color("green", "[EndPoint]"), l.endPoint)
		}
		if data != nil {
			l.log.Println(utils.Color("yellow", "[Data]"), data)
		}
	} else {
		l.log.Println(utils.Color("red", "[INFO]"), strings.Title(data.(error).Error()))
		l.log.Println(utils.Color("green", "[EndPoint]"), l.endPoint)
	}
}
