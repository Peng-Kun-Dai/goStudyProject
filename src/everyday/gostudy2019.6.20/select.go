package main

import "fmt"

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	select {
	case u := <-ch1:
		fmt.Println(u)
	case n := <-ch2:
		fmt.Println(n)

	}
	switch <-ch1 {
	case 1:
		fallthrough
	case 2:
		fallthrough
	default:

	}
}
