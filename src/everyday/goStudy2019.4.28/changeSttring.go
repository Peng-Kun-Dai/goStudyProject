package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Hello 世界！"

	b := []byte(s) // 转换为 []byte，自动复制数据

	b[5] = ',' // 修改 []byte

	fmt.Printf("%s\n", s) // s 不能被修改，内容保持不变

	fmt.Printf("%s\n", b) // 修改后的数据
	r := []int32(s)
	r[5] = ','
	r[6] = '中'
	r[7] = '国'
	fmt.Println(s)
	fmt.Println(string(r))

	fmt.Println(len("我爱你"))
	fmt.Println(utf8.RuneCountInString("我爱你"))

}
