package main

import "fmt"

/*求最大公约数*/
func f4(x, y int) int {
	for y != 0 {
		x, y = y, x%y
		fmt.Println(x, " ", y)
	}
	return x
}

/*fibonacci*/
func f5(n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
}
func main() {
	//fmt.Println(f4(12, 18))
	f5(5)
}
