package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logName := "./filetmp/test.log"
	debug(logName)
}
func debug(logname string) {
	//打开文件（如果没有就创建）
	logFile, err := os.OpenFile(logname, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create %s err :%v", logname, err)
	}
	defer logFile.Close()
	//文件句柄、前缀，日志属性
	//debugLog是一个logger
	debugLog := log.New(logFile, "[Debug]", log.Ldate)
	debugLog.SetPrefix("[Debug]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.SetFlags(log.Ltime)
	debugLog.Println("this is a Debug log")
	debugLog.Println(rangearray())
}

func rangearray() (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	panic("An error occurred")
}
