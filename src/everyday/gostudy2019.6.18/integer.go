package main

import "fmt"

/*
	8bit			16bit	32bit			64bit
	int8			int16	int32（rune）	int64  有符号
	uint8（byte）	uint16	uint32			uint64 无符号

	对应特定cpu的int和uint	32或者64bit
	uintptr没有具体大小但是足以容纳指针
*/
/*
	算术运算符
	+	-	*	/	++	--
	关系运算符
	==	!=	>	<	>=	<=
	逻辑运算符
	&&	||	!
	位运算符
	&	|	^	<<	>>
	赋值运算符
*/
func main() {
	fmt.Println(5 / 4)     //1
	fmt.Println(5.0 / 4.0) //1.25
	fmt.Println(5 / 4.0)   //1.25
	fmt.Println(5.0 / 4)   //1.25
	var x complex64 = complex(1, 2)
	var y complex64 = complex(1, 2)
	fmt.Println(x * y)
}
