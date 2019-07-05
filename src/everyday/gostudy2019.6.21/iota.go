package main

import "fmt"

const (
	a = iota
	b
	c = 100
	d
	e = iota
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
