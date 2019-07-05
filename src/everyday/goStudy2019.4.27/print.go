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
	fmt.Printf("everyday %d\n", 23)
	fmt.Println(map[int]string{3: "everyday"})
	fmt.Printf("everyday %v\n", map[int]string{2: "everyday"})
	fmt.Fprint(os.Stdout, "everyday ", 23, "\n")
	//fmt.Println("everyday", 23)
	//fmt.Println(fmt.Sprint("everyday ", 23))
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	//fmt.Printf("%#v\n", timeTwo)
}
