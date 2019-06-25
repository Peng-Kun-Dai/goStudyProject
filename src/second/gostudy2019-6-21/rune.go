package main

import "fmt"

func main() {
	//var chinese []rune = make([]rune, 10)
	s := "你好"
	chinese := []rune(s)
	fmt.Println(s)
	chinese = append(chinese, '嗨')
	fmt.Println(string(chinese))
}
