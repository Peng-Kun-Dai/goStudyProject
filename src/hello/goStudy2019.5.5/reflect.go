package main

import (
	"fmt"
	"reflect"
)

func main() {

	type MyInt int
	var m MyInt = 5
	t := reflect.TypeOf(m)
	fmt.Println(t)
	fmt.Println(t.Kind())
	v := reflect.ValueOf(m)
	fmt.Println(v)
	fmt.Println(v.Kind())
	fmt.Println(v.Interface())
}
