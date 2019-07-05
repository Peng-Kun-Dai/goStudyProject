package main

//基本数据类型

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func zero() (int, float64, bool, string) {
	var i = 0
	var f float64
	var b bool
	var s = "everyday"
	return i, f, b, s
}

func main() {
	fmt.Printf("Type: %T     Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T     Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T     Value: %v\n", z, z)
	fmt.Println(real(z), imag(z))
	fmt.Println(zero())
}

/*
	bool

	string

	int  int8  int16  int32  int64  //有符号
	uint uint8 uint16 uint32 uint64 uintptr  //无符号

	byte // uint8 的别名

	rune // int32 的别名
    // 表示一个 Unicode 码点

	float32 float64

	complex64 complex128  32位和64位的复数
*/
