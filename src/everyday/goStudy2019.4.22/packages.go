package main

import (
	"fmt"
	"math/rand"
)

/*
	推荐使用组合导包
*/

func main() {
	fmt.Println("a number is", rand.Intn(100))
	fmt.Println("a number is", rand.Intn(100))
	fmt.Println("a number is", rand.Float64())
	fmt.Println("a number is", rand.NewSource(2))

}
