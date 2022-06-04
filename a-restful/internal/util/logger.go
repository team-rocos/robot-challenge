package util

import (
	"fmt"
	"log"
)

//type ILogger interface {
//	Info(format string, args ...interface{})
//	Debug(format string, args ...interface{})
//}

type Logger struct {
	name string
}

func NewLogger(name string) *Logger {
	return &Logger{
		name: name,
	}
}

func (l Logger) Info(format string, args ...interface{}) {
	log.Println(fmt.Sprintf(format, args...))

}

func (l Logger) Debug(format string, args ...interface{}) {
	l.Info(format, args...)
}
