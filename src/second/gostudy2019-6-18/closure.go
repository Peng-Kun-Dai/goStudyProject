package main

import "fmt"

type closefunc func()

func main() {
	//匿名函数
	closurefunc := func(x, y int) int {
		return x + y
	}
	fmt.Println(closurefunc(3, 5))   //通过变量名调用
	fmt.Println(func(x, y int) int { //直接使用
		return x + y
	}(3, 5))

}
