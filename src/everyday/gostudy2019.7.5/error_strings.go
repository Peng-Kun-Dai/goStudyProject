package main

import "bytes"

//误用字符串
func main() {
	var b bytes.Buffer
	for condition {
		b.WriteString(str) // 将字符串str写入缓存buffer
	}
	return b.String()
}

//由于编译优化和依赖于使用缓存操作的字符串大小，当循环次数大于15时，效率才会更佳。
