package main

import (
	"fmt"
	"math"
)

/*
在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。

当右值声明了类型时，新变量的类型与其相同：

不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度：
*/
func main() {
	v := 42.8 + 0.8i + 4 + math.Pi // 修改这里！
	//var v = 42.8 + 0.5i
	fmt.Printf("v is of type %T\n", v)
}
