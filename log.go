// main.go
package main

import (
	"fmt"
	"path/filepath"
	"os"
	"time"
)

const TimeFormatFile string = "2006-01-02T15-04-05"
const TimeLayoutYYYYMMDD_HHMMSS = "2006-01-02 15:04:05"

func logError(location string,err error){
	fmt.Println("error in location : " + location + " - " + err.Error())
	//TODO:
	//logFile := GetCurrentDir() + time.Now().Format(TimeFormatFile )
	//logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//Log.New...
}

func GetCurrentDir()  string {
	if ex, err := os.Executable();err == nil {
		return filepath.Dir(ex)
	}	//TODO : else ....
	return ""
}

func NowTimeStr() string {
	return time.Now().Format(TimeLayoutYYYYMMDD_HHMMSS)
}
