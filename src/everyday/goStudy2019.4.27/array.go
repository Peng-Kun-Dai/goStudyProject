package main

/*
数组是值。将一个数组分配给另一个数组会复制所有元素。
特别是，如果将数组传递给函数，它将接收数组的副本，而不是指向它的指针。
数组的大小是其类型的一部分。类型[10]int 和[20]int不同。
value属性可能有用，但也很昂贵; 如果你想要类似C的行为和效率，你可以传递一个指向数组的指针。=使用切片
*/
func main() {
	array := [...]float64{7.0, 8.5, 9.1}
	x := Sum(&array) //注意显式地址运算符
}
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}

	return sum
}
