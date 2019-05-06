package main

import "fmt"

/*
go 的for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的。
*/

func main() {
	fmt.Println(for1())
	fmt.Print(for2())

}

func for1() int {
	sum := 0
	for i := 0; i < 100; i++ {
		//sum = sum + i
		sum += i
	}
	return sum
}

//初始化语句和后置语句是可选的。
func for2() int {
	sum := 1
	for sum < 100 { //或者直接for sum<1000{}
		//sum = sum + sum
		sum += sum
	}
	return sum
}

//无限循环
func for3() {
	//报错
	/*for {
	}*/
}
