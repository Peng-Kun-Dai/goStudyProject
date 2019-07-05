package main

import (
	"fmt"
	"strings"
)

//切片可以理解为动态数组(不恰当)，数组的引用吧
/*
每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。
在实践中，切片比数组更常用。
类型 []T 表示一个元素类型为 T 的切片。
切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：
a[low : high]
它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：
a[1:4]  前闭后开
*/
func slice1() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:6] //primes[1；4)
	//s[1] = 4
	fmt.Print(s)
}

/*
切片就像数组的引用
切片并不存储任何数据，它只是描述了底层数组中的一段。
更改切片的元素会修改其底层数组中对应的元素。
与它共享底层数组的切片都会观测到这些修改。
*/
func slice2() {
	name := [4]string{
		"John",
		"Lucifer",
		"Jayce",
		"Alice",
	}
	var s1 = name[0:2]
	s2 := name[1:4]
	fmt.Println(name)
	fmt.Println(s1)
	fmt.Println(s2)
	s2[0] = "Bob" //这里S2的第一个元素是name的第二个元素
	fmt.Println(name)
	fmt.Println(s1)
	fmt.Println(s2)

}

func slice3() {
	q := []int{2, 3, 5, 7, 9}
	//q = []int{2, 5, 3, 8}
	fmt.Println(q)
	R := []bool{true, false, true, true, false, true}
	fmt.Println(R)
	s := []struct {
		i int
		b bool
	}{
		{1, false},
		{3, true},
		{3, false},
	}
	fmt.Println(s)
}

/*
在进行切片时，你可以利用它的默认行为来忽略上下界。
切片下界的默认值为 0，上界则是该切片的长度。
*/
func slice4() {
	s := []int{2, 3, 6, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
	s = s[:4]
	//var a = s[:6]
	fmt.Println(s)
}

/*
切片拥有 长度 和 容量。
切片的长度就是它所包含的元素个数。
切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
*/
func slice5() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:6]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//nil切片
func slice6() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("this slice is nil")
	}
}

/*
切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。
make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
要指定它的容量，需向 make 传入第三个参数：
*/
func slice7() {
	a := make([]int, 5)
	printSlice(a)
	b := make([]int, 0, 5)
	printSlice(b)
	c := b[:2]
	printSlice(c)
	//printSlice(a)
}

func slice8() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func slice9() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[0][1] = "X"
	board[0][2] = "O"
	board[1][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

//追加切片
func slice10() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}
func main() {
	slice10()
	//pic.Show(Pic)

}
func Pic(dx, dy int) [][]uint8 {

	var target [][]uint8
	var small []uint8
	var number uint8
	for i := 0; i < dy; i++ {

		small = small[:0]
		for j := 0; j < dx; j++ {
			number = uint8(i * j)
			small = append(small, number)
		}
		target = append(target, small)
	}
	return target
}
