package logger

import (
	"log"
	"os"
)

type any = interface{}

var (
	DebugLog   *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
	initLog    *log.Logger
)

func InitLogger() {
	DebugLog = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLog = log.New(os.Stderr, "INIT: ", log.Ldate|log.Ltime|log.Lshortfile)
	initLog.Println("Init api")
}

func Debugln(text ...any) {
	DebugLog.Println(text...)
}

func Debugf(text string, v ...any) {
	DebugLog.Printf(text, v...)
}

func Warningln(text ...any) {
	WarningLog.Println(text...)
}

func Warningf(text string, v ...any) {
	WarningLog.Printf(text, v...)
}

func Errorln(text ...any) {
	ErrorLog.Println(text...)
}

func Errorf(text string, v ...any) {
	ErrorLog.Printf(text, v...)
}

func Fatalln(text ...any) {
	ErrorLog.Println(text...)
}

func Fatalf(text string, v ...any) {
	ErrorLog.Fatalf(text, v...)
}
