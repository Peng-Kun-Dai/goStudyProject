package main

import "fmt"

//结构体内嵌   继承
//子类型的方法会变成外部类型的方法，但是接收参数的任然是子类型字段
type Eat struct {
	str1 string
}
type Drink struct {
	str2 string
}
type people struct {
	*Eat
	*Drink
	*Eatson
	str1 string
}
type Eatson struct {
	str1 string
}

func main() {
	eat := Eat{""}
	drink := Drink{""}
	eatson := Eatson{""}
	my := &people{&eat, &drink, &eatson, ""}

	//var my people  无法访问内部的类型，因为空指针异常
	my.str1 = "everyday"
	my.Eat.str1 = "Eat everyday"
	my.Eatson.str1 = "Eatson everyday"
	fmt.Println(my.str1)
	fmt.Println(my.Eat.str1)
	fmt.Println(my.Eatson.str1)
}
