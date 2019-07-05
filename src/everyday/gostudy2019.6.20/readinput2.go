package main

//使用 fallthrough 会强制执行后面的 case 语句，
// fallthrough 不会判断下一条 case 的表达式结果是否为 true。
import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var inputstring string
var err error

func main() {
	//创建读取器
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")

	inputstring, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", inputstring)
	}
}
