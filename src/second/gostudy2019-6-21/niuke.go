package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 0)
	s = append(s, 1)
	fmt.Println(s)
	x := []int{
		1, 2, 3,
		4, 5, 6}
}
