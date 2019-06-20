package test

import (
	"fmt"
	"log"
	"os"
)

func main() {

	/*ch := make(chan int)
	go test(ch)
	<-ch*/
	//test_again()

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

	//fmt.Println("hello")
	//os.Args 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	//你可以使用标准的索引位置方式取得单个参数的值。
	//PS C:\go-second\src\awesomeProject\src\hello\test> ./test a s d f g h
	//[C:\go-second\src\awesomeProject\src\hello\test\test.exe a s d f g h]
	//[a s d f g h]
	//d
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

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
