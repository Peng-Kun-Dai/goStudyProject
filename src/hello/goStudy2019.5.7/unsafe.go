package main

import (
	"fmt"
	"unsafe"
)

const (
	a = "hello world"
	b = len(a)
	c = unsafe.Sizeof(a)
	d = unsafe.Sizeof(b)
)

func main() {
	fmt.Println(a, b, c, d)
}
