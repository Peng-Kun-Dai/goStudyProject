package main

import "fmt"

func main() {
	a := 10
	add(a)
	fmt.Println(a) //10
	add2(&a)
	fmt.Println(a)
	pop()
}

//值传递
func add(a int) int {
	a++
	return a
}

//传递指针的消耗更低
//引用传递
func add2(a *int) *int {
	*a++
	return a
}
func pop() {
	fmt.Println("everyday")
}
