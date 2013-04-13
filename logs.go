package logs

import (
	"fmt"
	"log"
	"os"
)

type Level int

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

var CurrentLevel Level
var Logger *log.Logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

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
		case FATAL:
			str = "FATAL"
		default:
			panic("what level?")
		}
		first, isString := v[0].(string)
		remaining := v[1:]
		if isString {
			output := fmt.Sprintf("["+str+"] "+first, remaining...)
			Logger.Output(3, output)
		} else {
			slice := []interface{}{"[" + str + "]", first}
			slice = append(slice, remaining...)
			output := fmt.Sprintln(slice...)
			Logger.Output(3, output)
		}
	}
}

func Debug(v ...interface{}) {
	Log(DEBUG, v...)
}

func Info(v ...interface{}) {
	Log(INFO, v...)
}

func Warn(v ...interface{}) {
	Log(WARN, v...)
}

func Error(v ...interface{}) {
	Log(ERROR, v...)
}

func Fatal(v ...interface{}) {
	Log(FATAL, v...)
	os.Exit(1)
}
