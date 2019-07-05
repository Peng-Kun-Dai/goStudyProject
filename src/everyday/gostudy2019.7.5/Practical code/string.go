package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//如何修改字符串中的一个字符：
	str := "everyday"
	c := []byte(str)
	c[0] = 'H'
	s2 := string(c)
	fmt.Println(s2)
	//获取字符串的字串
	//substr:=str[n:m] //下标
	//如何使用for或者for-range遍历一个字符串
	for i := 0; i < len(str); i++ {
		_ = str[i]
	}
	for i, v := range str {
		println(i, v)
	}
	//如何获取一个字符串的字节数以及字符数
	str2 := "中国"
	println(len(str2))
	println(utf8.RuneCountInString(str2))
	println(len([]rune(str2)))
	//如何连接字符串：
	//bytes.Buffer
	//strings.Join()
	//+=
}
