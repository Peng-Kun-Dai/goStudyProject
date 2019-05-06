package main

import (
	"bufio"
	"fmt"
	"io"
	_ "log"
	"os"
)

func main() {
	//writefile()
	copyfile()
}
func writefile() {
	//OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。
	// 它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，
	// 返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
	file4, err := os.OpenFile("./filetmp/file4.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	defer file4.Close()
	//写入器变量
	writer := bufio.NewWriter(file4)
	str1 := "hello go"
	for i := 0; i < 3; i++ {
		_, _ = writer.WriteString(str1)
	}
	//Flush方法将缓冲中的数据写入下层的io.Writer接口。
	_ = writer.Flush()
}
func copyfile() {
	src, err := os.Open("./filetmp/file4.txt")
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile("./filetmp/file4_copy.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
}
