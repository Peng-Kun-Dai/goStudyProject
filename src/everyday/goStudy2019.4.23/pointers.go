package main

import "fmt"

/*
	指针*/
func main() {
	//i, j := 32, 42
	i := 32
	j := 42

	p := &i //简单赋值语句，同时指明了p是一个int型指针
	fmt.Println(*p)
	*p = 33
	fmt.Println(i)
	fmt.Println(p)

	p = &j //若J不为int型则报错
	*p = *p / 7
	fmt.Println(j)
	fmt.Println(p)
}
