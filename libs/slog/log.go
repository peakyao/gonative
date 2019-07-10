package slog

import (
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

type LogStruct struct {
	Logpath string
	Logfile string
}

func New() *LogStruct {
	return &LogStruct{}
}

/**
 * 信息日志
 **/
func (l *LogStruct) InfoLog(str interface{}) {
	infofile, err := os.OpenFile("logs/info.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Info = log.New(io.MultiWriter(infofile, os.Stderr), "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	Info.Println(str)
}

/**
 * 错误日志
 **/
func (l *LogStruct) ErrorLog(str interface{}) {

	errorfile, err := os.OpenFile("logs/errors.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Error = log.New(io.MultiWriter(errorfile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error.Panicln(str)

}

func InfoLog(str ...interface{}) {
	infofile, err := os.OpenFile("logs/info.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Info = log.New(io.MultiWriter(infofile, os.Stderr), "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	Info.Println(str)
}

func ErrorLog(str interface{}) {

	errorfile, err := os.OpenFile("logs/errors.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Error = log.New(io.MultiWriter(errorfile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error.Panicln(str)

}

/**
 * 监控函数日志
 **/
func FuncLog(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		InfoLog("Handler function called - " + name)
		h(w, r)
	}
}
