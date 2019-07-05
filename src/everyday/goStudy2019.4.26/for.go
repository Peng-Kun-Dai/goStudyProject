package main

import "fmt"

func main() {

	//类似Java
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
	fmt.Printf("\nfor1 over")

	//类似于while
	var condition int = 5
	for condition > 0 {
		condition--
		fmt.Print(condition, "  ")
	}
	fmt.Printf("\nfor2 over")

	//无限
	//使用break lable 和 goto lable 都能跳出for循环；
	// 不同之处在于：break标签只能用于for循环，且标签位于for循环前面，
	// goto是指跳转到指定标签处
	//跳出for的两种方法break，goto
Loop:
	for {
		condition++
		fmt.Print(condition, "  ")
		if condition > 10 {
			//goto Loop
			//直接break也行
			break Loop
		}
	}
	//Loop:
	fmt.Println("for3 over")
}
