package main

import "fmt"

/*
Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的名称应当具有一定的意义，它可以作为文档使用。

没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。

直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
*/

func main() {
	fmt.Print(spilt(15))
}

func spilt(sum int) (x, y int) {
	x = sum + 2
	y = x + 2
	return
}
