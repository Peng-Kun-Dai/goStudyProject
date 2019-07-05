package main

import (
	"fmt"
	"time"
)

func main() {
	//test of time
	start := time.Now()
	defer func() {
		end := time.Now()
		timer := end.Sub(start)
		fmt.Println(timer)
	}()

	res := uint64(0)
	for i := 0; i <= 40; i++ {
		res = fibonacciOfCache(i)
		//res=fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, res)
	}
}

//basic version
func fibonacci(n int) (result uint64) {

	if n <= 1 {
		result = 1
	} else {
		result = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

//cache version
/*斐波那契数列内存缓存版本实现*/
const lim = 41

//缓存
var fibs [lim]uint64

func fibonacciOfCache(n int) (result uint64) {

	//If it is already in the cache, take it out directly
	if fibs[n] != 0 {
		result = fibs[n]
		return
	}
	if n <= 1 {
		result = 1
	} else {
		result = fibonacciOfCache(n-1) + fibonacciOfCache(n-2)
	}
	fibs[n] = result //put into cache
	return
}
