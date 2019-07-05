package main

import (
	"fmt"
	"os"
)

/*i++是语句，不是表达式，不能用于赋值
++和--只能在变量之后*/
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
