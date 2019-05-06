package main

import (
	"fmt"
	"os"
)

//打印输出
//格式  %d  对应数字格式
//		%v  通用格式
type T struct {
	a int
	b float64
	c string
}

func main() {
	fmt.Printf("hello %d\n", 23)
	fmt.Println(map[int]string{3: "hello"})
	fmt.Printf("hello %v\n", map[int]string{2: "hello"})
	fmt.Fprint(os.Stdout, "hello ", 23, "\n")
	//fmt.Println("hello", 23)
	//fmt.Println(fmt.Sprint("hello ", 23))
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	//fmt.Printf("%#v\n", timeTwo)
}
