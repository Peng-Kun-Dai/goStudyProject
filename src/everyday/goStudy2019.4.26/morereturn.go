package main

import "fmt"

//函数返回值
func nextInt() (int, int) {
	return 1, 2
}

func namereturn() (a, b int) {
	//a,b会被默认初始化
	//a, b = 3, 4
	return
}

func main() {
	fmt.Print(namereturn())
}
