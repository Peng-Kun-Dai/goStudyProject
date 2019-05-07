package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	//func New() hash.Hash
	//返回一个新的使用SHA1校验的hash.Hash接口。
	hasher := sha1.New()
	//func WriteString(w Writer, s string) (n int, err error)
	//WriteString函数将字符串s的内容写入w中。如果w已经实现了WriteString方法，函数会直接调用该方法。
	_, _ = io.WriteString(hasher, "test")
	fmt.Println(hasher)
	b := []byte{}
	//const Size = 20
	//SHA1校验和的字节数。
	//func Sum(data []byte) [Size]byte
	//返回数据data的SHA1校验和。
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
	fmt.Printf("Result: %o\n", hasher.Sum(b))
	// Reset resets the Hash to its initial state.
	hasher.Reset()
	fmt.Println(hasher)
	fmt.Println("-------------------------")
	data := []byte("We shall overcome")
	//Write (via the embedded io.Writer interface) adds more data to the running hash.
	// It never returns an error.
	n, err := hasher.Write(data)
	fmt.Println(hasher)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
