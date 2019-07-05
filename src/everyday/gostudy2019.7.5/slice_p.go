package main

//不需要将一个指向切片的指针传递给函数
//切片实际是一个指向潜在数组的指针。
// 我们常常需要把切片作为一个参数传递给函数是因为：实际就是传递一个指向变量的指针，
// 在函数内可以改变这个变量，而不是传递数据的拷贝。
//因此应该这样做：

//`func findBiggest( listOfNumbers []int ) int {}`
//而不是：

//`func findBiggest( listOfNumbers *[]int ) int {}`
