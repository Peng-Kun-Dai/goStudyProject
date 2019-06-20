package main

import "fmt"

const c1 = 3

/*const c2  = hello()*/ //常量的值必须在编译期就能够确定

const c3 = 123456

func hello() int {
	return 3
}
func main() {
	fmt.Println(c3)
}
