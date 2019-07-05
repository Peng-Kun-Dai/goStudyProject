package main

import (
	"fmt"
	"log"
)

//函数产生下标越级异常对于一些隐式的运行时错误，如切片索引越界、类型断言错误等情形下，
// panic方法就会被调用,它将立刻中断当前函数的执行，并展开当前Goroutine的调用栈，
// 依次执行之前注册的defer函数。
// 当栈展开操作达到该Goroutine栈顶端时，程序将终止。
// 但这时仍然可以使用Go的内建recover方    法重新获得Goroutine的控制权，并将程序恢复到正常执行的状态。
//调用recover方法会终止栈展开操作并返回之前传递给panic方法的那个参数。
// 由于在栈展开过程中，只有defer型函数会被执行，因此recover的调用必须置于defer函数内才有效。

//错误恢复代码会对返回的错误类型进行类型断言，判断其是否属于Error类型。
// 如果类型断言失败，则会引发运行时错误，并继续进行栈展开，最后终止程序 ,这个过程将不再会被中断。
// 类型检查失败可能意味着程序中还有其他部分触发了panic，如果某处存在索引越界访问等，
// 因此，即使我们已经使用了panic和recover机制来处理解析错误，程序依然会异常终止。
func safedo() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("函数计算失败", err)
		}
	}()
	var array [3]int
	var s = make([]int, 0)
	s = array[:]
	fmt.Println(s[7])
}
func main() {
	safedo()
	//第二次的panic并没有解决
	var array [3]int
	var s = make([]int, 0)
	s = array[:]
	fmt.Println(s[7])
}
