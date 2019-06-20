package main

import "fmt"

var arr1 [5]int        //arr1是值类型
var arr2 = new([5]int) //是*【5】int

func main() {

	arr3 := arr2
	for i := range arr3 {
		arr3[i] = i
	}
	for _, sum := range arr2 {
		fmt.Println(sum)
	}

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	var arr4 [3][5]int //[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
	fmt.Println(arr4)
}
