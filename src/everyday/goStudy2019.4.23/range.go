package main

import "fmt"

var powslice = []int{1, 2, 4, 8, 16, 32, 64}

func main() {
	range2()
}
func range1() {
	for i, v := range powslice {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
func range2() {
	powslices := make([]int, 10)
	for i := range powslices {
		powslices[i] = 1 << uint(i) //==2**i
	}
	for _, value := range powslices {
		fmt.Printf("%d\n", value)
	}
}
