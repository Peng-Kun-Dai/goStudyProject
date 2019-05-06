package main

//导出名
import (
	"fmt"
	"math"
)

func main() {
	/*
		大写开头代表已导出名，可外部以使用，
		小写则相反，只能包内使用
	*/
	//fmt.Println(math.pi)
	fmt.Print(math.Pi)
}
