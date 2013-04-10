package logs

import (
	"log"
)

type Level int

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var CurrentLevel Level

func SetLevel(level Level) {
	CurrentLevel = level
}

func Log(level Level, v ...interface{}) {
	if level >= CurrentLevel {
		str := ""
		switch level {
		case DEBUG:
			str = "DEBUG"
		case INFO:
			str = "INFO"
		case WARN:
			str = "WARN"
		case ERROR:
			str = "ERROR"
		default:
			panic("what level?")
		}
		first := v[0].(string)
		remaining := v[1:]
		log.Printf("["+str+"]"+first, remaining...)
	}
}

func Debug(v ...interface{}) {
	Log(DEBUG, v)
}

func Info(v ...interface{}) {
	Log(INFO, v)
}

func Warn(v ...interface{}) {
	Log(WARN, v)
}

func Error(v ...interface{}) {
	Log(ERROR, v)
}
