package main

func main() {
	//创建
	//arr1:=new([len]type)  数组
	//slice1:=make([]type,len)  切片
	//初始化：
	//arr1 := [...]type{i1, i2, i3, i4, i5}
	//arrKeyValue := [len]type{i1: val1, i2: val2}
	//var slice1 []type = arr1[start:end]
	//如何截断数组或者切片的最后一个元素：
	//line = line[:len(line)-1]
}
func find() {
	//如何在一个二维数组或者切片arr2Dim中查找一个指定值V：

	found := false
Found:
	for row := range arr2Dim {
		for column := range arr2Dim[row] {
			if arr2Dim[row][column] == V {
				found = true
				break Found
			}
		}
	}
}
