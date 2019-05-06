package main

//导包，接口
import (
	"awesomeProject/src/hello/test"
	"fmt"
)

type methoder interface {
	method()
}
type mInt int

func (i mInt) method() {
	fmt.Println("value is ", i)
}

func main() {
	//fmt.Println(8<<1 + 5<<2)
	var i mInt = 5
	i.method()
	var stu1two test.Stu1
	fmt.Println(stu1two)
	stu1two.Setstu("luci", 12)
	fmt.Println(stu1two)
	fmt.Println(stu1two.Ostu())
}
