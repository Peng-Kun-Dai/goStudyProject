package test

import (
	"fmt"
	"time"
)

func main() {
	b()
}
func trace(s string) { fmt.Println("entering:", s, time.Now().UTC()) }

func untrace(s string) { fmt.Println("leaving:", s, time.Now().UTC()) }

func a0() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	time.Sleep(time.Second)
	defer untrace("b")
	fmt.Println("in b")
	a0()
}
