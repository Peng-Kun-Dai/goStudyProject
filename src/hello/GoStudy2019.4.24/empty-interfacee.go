package main

import "fmt"

//空接口
/*
指定了零个方法的接口值被称为 *空接口：*
interface{}
空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
*/
func main() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)

	var f float64
	i = f
	describe2(i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
