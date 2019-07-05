package main

import "fmt"

/*
var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。

就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。
在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。

函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用(没有关键字)。
*/
var type1, type2, type3 bool //默认false

func main() {
	var i int = 1 //
	k := 3
	fmt.Print(i, k, type1, type2, type3)
}
