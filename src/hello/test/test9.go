package main

import "fmt"

func main() {
	/*v := tt4()
	fmt.Println(v())*/
	fmt.Println(tt4()())
}
func tt4() func() int {
	var i = 0
	defer func() {
		i++
	}()
	i++
	return func() int {
		return i
	}
}
