package main

import "fmt"

/*
	多值返回
*/

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("everyday", "world")
	fmt.Println(a, b)
}
