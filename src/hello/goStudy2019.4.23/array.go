package main

import "fmt"

/*
数组的长度是其类型的一部分，因此数组不能改变大小
*/
func main() {

	var a [2]string //长度为2的字符串数组
	a[0] = "jayce"
	a[1] = "lucifer"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	intarray := [3]int{4, 1, 8}
	fmt.Println(intarray)
}
