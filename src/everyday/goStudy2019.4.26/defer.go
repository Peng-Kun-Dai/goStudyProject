package main

import (
	"fmt"
	"io"
	"os"
)

/*
互斥锁，文件开关
defer执行压栈时参数会求值，而不是在出栈调用执行时，变量的值不会随着函数的执行而改变
*/
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		//代表发生了错误
		return "", err
	}
	defer f.Close() //f.close will run when we're finished

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:1])
		result = append(result, buf[0:n]...)
		if err != io.EOF {
			break
		}
		return "", err
	}
	return string(result), nil

}
func main() {
	bb()
}
func trace(s string) string {
	fmt.Println("entering", s)
	return s
}
func un(s string) {
	fmt.Println("leaveing:", s)
}
func aa() {
	defer un(trace("aa"))
	fmt.Println("in aa")
}
func bb() {
	defer un(trace("bb"))
	fmt.Println("in bb")
	aa()
}
