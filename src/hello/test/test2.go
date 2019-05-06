package main

import "fmt"

func main() {
	var str = "我爱你，中国"
	for _, v := range str {
		fmt.Println(string(v))
	}
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}
