package main

import (
	"fmt"
	"log"
)

func main() {

	/*ch := make(chan int)
	go test(ch)
	<-ch*/
	test_again()

	//main发生panic
	/*defer func() {
		if err := recover(); err != nil {
			log.Println("run time panic :%v", err)
		}
	}()
	panic("手动触发panic")
	//panic("第二次触发panic")
	fmt.Println("我不会被执行")
	*/

	fmt.Println("hello")

}

//go  test()
func test(ch chan int) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("run time panic :%v", err)
		}
		ch <- 1
	}()
	panic("手动触发panic")
	//panic("第二次触发panic")
	fmt.Println("我不会被执行")
}

//直接func()
func test_again() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("run time panic :%v", err)
		}
	}()
	panic("手动触发panic")
	//panic("第二次触发panic")
	fmt.Println("我不会被执行")

}
