package main

import "fmt"

func main() {
	var a = 10
	add(&a)
	fmt.Println(a)

}
func add(p *int) {
	*p++
}
