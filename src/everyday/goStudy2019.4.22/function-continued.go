package main

import "fmt"

/*
	函数有多个已命名形参类型相同时，除最后一个类型以外，其它都可以省略
*/

func add2(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println(add2(1, 2, 3))
}
