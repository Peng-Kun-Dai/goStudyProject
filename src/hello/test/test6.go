package test

import "fmt"

func main() {
	//var array = []int{1, 4, 6, 7, 4}
	//fmt.Println(min(array...))
	//printtest(array...)
	a()
}
func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
func printtest(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
