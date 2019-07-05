package main

type name interface {
	method1(a int) int
	method2(a int) int
}
type inter int

func main() {
	var ai name
	var s inter = 10
	ai = s
	ai.method1(int(s))

}

func (i inter) method1(a int) int {
	a++
	return a
}

func (i inter) method2(a int) int {
	a--
	return a
}
