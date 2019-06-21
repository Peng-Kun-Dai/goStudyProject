package main

import "fmt"

type intt int

func main() {
	var i intt = 10
	i++
	fmt.Println(i)
}

type struct1 struct {
	name string
	age  int
}
