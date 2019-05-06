package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
 	Go 只运行选定的 case，而非之后所有的 case。 实际上，
	Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。
	除非以 fallthrough 语句结束，否则分支会自动终止。
	Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数
*/

func main() {
	switch2()
	switch3()
}
func switch1() {
	fmt.Print("Go runs on ")
	/*os := runtime.GOOS
	fmt.Println(os)
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}*/
	//或者
	//变量os定义的位置决定它的作用域不一样
	//fmt.Println(os)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

//switch 的求值顺序
func switch2() {
	fmt.Println("when's Saturday?")
	today := time.Now().Weekday()
	switch time.Thursday {
	case today:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case 2 + today:
		fmt.Println("in two days")
	default:
		fmt.Println("Too far away")
	}
}

//没有条件的switch
func switch3() {
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("good morning")
	case time.Hour() < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good evening")

	}
}
