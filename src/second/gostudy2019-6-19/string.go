package main

import (
	"fmt"
)

func main() {
	str := "你好中国"
	b := []byte(str)
	fmt.Println(b)
	r := []rune(str)
	fmt.Println(r)
}
