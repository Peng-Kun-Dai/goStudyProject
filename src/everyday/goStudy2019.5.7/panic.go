package main

import (
	"fmt"
	"log"
)

func main() {
	defer fmt.Println("enen")
	defer func() {
		fmt.Println("我应该被执行了")
		if err := recover(); err != nil {
			log.Println(err) //defer的后续代码会执行
			//log.Fatal(err)  //defer的后续代码不会被执行
		}
		fmt.Print("我在defer中，但是在recover之后，我能被执行吗？")
	}()
	defer fmt.Println("haha")
	panic("i want you runtime error")
	fmt.Println("我肯定不能被执行")
}
