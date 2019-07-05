package main

import (
	"fmt"
	"unsafe"
)

//字符串连接
func stringcontent() {
	str := fmt.Sprintf("asd%d", 123)
	fmt.Println(str)

	str1 := "asd" + "123"
	fmt.Println(str1)

}

//for循环
//当次循环continue语句之后的语句不再执行
//break 默认推出整个for循环

func forfunc() {
	//Loop:
	for i := 0; i < 4; i++ {
		for j := 0; j <= 4; j++ {
			fmt.Print("", j, "  ")
			if j == 2 {
				//break Loop  跳出整个Loop点
				//goto Loop  跳到Loop点继续执行
			}
			fmt.Print("--")

		}
		//Loop:
		fmt.Println()
	}
	//Loop:
}

//init 函数不能被调用
func init() {
}

//函数调用
func add(args ...int) int { //代表接受多个不定数量int型参数传入
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

//类型转化
/*type MyInt
var i int =1
var jMyInt  =(MyInt)i

*/
//常量定义
const zero = "A" //未使用的常量不会引起编译报错
const (
	s = "everyday"
	b = len(s)
	c = unsafe.Sizeof(int8(1))
	a //a="everyday"    不提供a的类型和初始化，a就等于它的上一个常量
)

func init() {
	//var m int

}
func main() {
	//forfunc()
	//var a = [3]int{1, 2, 3}
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 3, 5))
	//fmt.Println(add(a))			//
	//fmt.Println(add([]int{1,3}))    //
	fmt.Println(add([]int{1, 3, 4}...)) //[]byte{}...代表切片被打散传入

	//var _i = zero
	//i:=10
	//var i int = 3.4    //编译无法通过
	//i=10
	fmt.Println(a)
	//rune或者叫int32这种类型没有len方法
	//一个汉字三个字节
	fmt.Println(len(string(int32('你'))))

}
