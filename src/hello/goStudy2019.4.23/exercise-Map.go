package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	var m map[string]int
	m = make(map[string]int)
	var slice []string = strings.Fields(s)

	for i := 0; i < len(slice); i++ {
		_, ok := m[slice[i]]
		if ok {
			m[slice[i]]++
		} else {
			m[slice[i]] = 1
		}
	}
	return m
	//return map[string]int{"x": 1}
}

func main() {
	fmt.Println(WordCount("i love you ,and you"))
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}
