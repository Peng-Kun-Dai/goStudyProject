package main

import "fmt"

/*
创建新的类型名称
创建的新类型与其底层类型同样无法兼容
首字母大写包外可见
*/

type celsius float64    //摄氏温度
type fahrenheit float64 //华氏温度
type point *int

const (
	absoluteZeroC celsius = -273.15 //绝对零度
	freezingC     celsius = 0       //冰点
	boilingC      celsius = 100     //沸点
)

func main() {
	var sum float64 = 32.0
	var num celsius
	num = celsius(sum)
	fmt.Println(num)
	a := 10
	p := &a
	var pi point = (*int)(p)
	fmt.Println(pi)
}
