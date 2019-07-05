package main

import "os"

// 发生错误时使用defer关闭一个文件

func main() {
	for _, file := range files {
		if f, err = os.Open(file); err != nil {
			return
		}
		// 这是错误的方式，当循环结束时文件没有关闭
		defer f.Close()
		// 对文件进行操作
		f.Process(data)
	}

	for _, file := range files {
		if f, err = os.Open(file); err != nil {
			return
		}
		// 对文件进行操作
		f.Process(data)
		// 关闭文件
		f.Close()
	}
}
