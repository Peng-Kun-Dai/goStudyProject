package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	i, ok := strconv.ParseInt("-1000", 10, 0)
	if ok == nil {
		fmt.Printf("ParseInt , i is %v , type is %v\n", i, reflect.TypeOf(i))
	}

	ui, ok := strconv.ParseUint("100", 10, 0)
	if ok == nil {
		fmt.Printf("ParseUint , ui is %v , type is %v\n", ui, reflect.TypeOf(i))
	}

	oi, ok := strconv.Atoi("100")
	if ok == nil {
		fmt.Printf("Atoi , oi is %v , type is %v\n", oi, reflect.TypeOf(i))
	}

}
