package main

import "fmt"

func modifySlice() {
	//v := make([]int, 4, 6)
	v := []int{1, 2, 3, 4}
	var v2 []int = v
	//cXZN
	for i := range v {
		v2 = append(v2, i) //

		fmt.Printf("Modify Slice: value:%v\n", v2)

	}
}
func modifyMap() {
	data := map[string]string{"a": "A", "b": "B", "c": "C"}

	var data2 map[string]string = data //data2和data指向的是同一个内存地址
	for k, v := range data {
		//K,V 互换
		data2[v] = k
		//delete(data, k)
		fmt.Println("modify Mapping", data2)
	}
}

func main() {
	//modifySlice()
	modifyMap()

}
func cnz() {
	fmt.Println("cnzinco")

}

//slice和map的类型不同
//指针
