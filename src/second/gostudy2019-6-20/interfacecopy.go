package main

import "fmt"

type name3 interface {
	method1()
}

type name4 interface {
	method2()
}
type sum struct {
	a int
}

func (h sum) method1() {
	h.a++
	fmt.Println(h.a)
}
func (h sum) method2() { //
	h.a--
	fmt.Println(h.a)
}
func main() {
	var h *sum = new(sum)
	h.a = 0
	var inter1 name3 = h
	var inter2 name4 = h
	(inter1).method1() //解引用发生在调用之前，method方法执行时拿到的是已经解引用的数据
	inter2.method2()
}
