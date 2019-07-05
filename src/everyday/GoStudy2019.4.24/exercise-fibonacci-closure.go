package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func(int) int {
	small := 0
	big := 1
	i := 0
	return func(x int) int {
		switch x {
		case 0:
			return small
		case 1:
			return big
		default:
			i = small + big
			small = big
			big = i
			return i

		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
