package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Jayce", 23}
	z := Person{"vaen", 200}

	fmt.Println(a, z)
}
