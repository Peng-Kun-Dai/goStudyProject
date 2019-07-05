package main

import "fmt"

func test() interface{} {
	if true {
		fmt.Println("i am true")
		return nil
	}
	//下面的均没有执行
	defer fmt.Println("i am defer")
	return 0
}
func main() {
	test()
}
