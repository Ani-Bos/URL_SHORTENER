package logger

import (
	"io"
	"log"
)

type Logger struct {
	L *log.Logger
}

func New(out io.Writer, prefix string, flags int) *Logger {
	return &Logger{L: log.New(out, prefix, flags)}
}

func (l *Logger) LogMessage(message string) {
	l.L.Println(message)
}

func (l *Logger) LogFatalMessage(msg string) {
	l.L.Fatal(msg)
}
