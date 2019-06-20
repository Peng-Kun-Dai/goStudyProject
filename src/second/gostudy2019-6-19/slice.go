package main

import "fmt"

func main() {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
		fmt.Println(item)
	}
}
func test1() {
	var arr = [5]int{0, 1, 2, 3, 4}
	slice1 := &arr
	slice2 := arr[:]
	slice3 := arr[0 : len(arr)-1]
	slice4 := arr[1:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	slice5 := make([]int, 5, 10)
	fmt.Println(cap(slice5))
}
