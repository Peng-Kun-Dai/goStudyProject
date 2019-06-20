package main

import "fmt"

func main() {
	v := 1
	f3(&v)
}
func f3(p *int) {
	fmt.Println(p)
	*p++
	fmt.Println(p)
}
func f2() *int {
	v := 1
	return &v
}
func f1() {
	x := 1
	p := &x //p,of type *int,point to x
	fmt.Println(p)
	fmt.Println(*p)
	*p++
}
