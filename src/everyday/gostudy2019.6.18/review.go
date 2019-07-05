package main

import (
	"fmt"
	"os"
)

var (
	HOME = os.Getenv("JAVA_HOME")
	//os.Getenv  获取环境变量的值
	USER   = os.Getenv("GOOS")
	GOROOT = os.Getenv("GOROOT")
)

func main() {
	//a:=float64(10)
	//fmt.Println(HOME, USER, GOROOT)
	n()
	m()
	n()
}

var a = "G"

func n() { print(a) }

func m() {
	a := "O"
	print(a)
	for {
		break
	}
	for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
	}
}
