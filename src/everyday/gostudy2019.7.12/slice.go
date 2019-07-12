package main

import "fmt"

func main() {
	var s []int = make([]int, 3, 10)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s[11])
}
