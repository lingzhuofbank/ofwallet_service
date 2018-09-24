package utils

import (
	"os"
	"github.com/ofwallet_service/utils/go-logging"
	"runtime"
)

var (
	Logger *logging.Logger = nil
)

func CreateLogger() {
	logger := logging.MustGetLogger("logger")
	//logging.SetLevel(logging.WARNING, "logger")
	//
	format := logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} %{shortfile} %{longfunc} >>> %{level:.4s} %{id:04d} %{message}%{color:reset}`,
	)

	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	logging.SetBackend(backend2Formatter)
	//logging.SetLevel(logging.INFO, "logger")

	Logger = logger

}


func ErrorLogger(message interface{}){
	pc,_,line,_:=runtime.Caller(1)
	Logger.Error(message,"function: ",runtime.FuncForPC(pc).Name(),"line: ",line)
}

func NoticeLooger(message string){
	pc,_,line,_:=runtime.Caller(1)
	Logger.Notice(message,"function: ",runtime.FuncForPC(pc).Name(),"line: ",line)
}