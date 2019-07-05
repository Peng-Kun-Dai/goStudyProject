package main

import "fmt"

/*
	传入参数时变量名在类型之前
*/

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 3))
}
