package main

import (
	"fmt"
	"math"
)

type shaper interface {
	Area() float32
}

type square struct {
	side float32
}

type circle struct {
	radius float32
}

func (sq *square) Area() float32 {
	return sq.side * sq.side
}

func (ci *circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf shaper
	sq1 := new(square)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
