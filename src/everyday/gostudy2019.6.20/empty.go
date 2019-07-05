package main

type name1 interface {
	method1()
}
type name2 interface {
	method1()
}
type hello struct {
	a int
}

func (h *hello) method1() {

}

func main() {
	var h *hello = new(hello)
	var inter1 name1 = h
	var inter2 name2 = h
}
