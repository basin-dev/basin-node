package log

import (
	"log"
	"os"
	"runtime"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
)

func InitLoggers() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
	Info = log.New(os.Stdout, Cyan+"INFO: ", log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, Yellow+"WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, Red+"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(os.Stderr, Purple+"DEBUG: ", log.Lshortfile)
}
