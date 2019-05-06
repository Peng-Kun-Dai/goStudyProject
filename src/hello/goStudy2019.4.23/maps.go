package main

/*
映射将键映射到值。
映射的零值为 nil 。nil 映射既没有键，也不能添加键。
make 函数会返回给定类型的映射，
文法（literals）枚举出来
*/
import "fmt"

type Vertex struct {
	Lat, Long float64
}

func maps1() {
	//声明  map[KeyType]ValueType
	var m map[string]Vertex
	fmt.Println(m)
	//初始化
	m = make(map[string]Vertex)
	fmt.Println(m)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}

func main() {
	var n map[int]string = map[int]string{12: "daipengkun", 13: "wanglei"}
	var m = map[int]string{12: "daipengkun", 13: "wanglei"}
	var mn = map[string]Vertex{
		"Jayce": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(n)
	//fmt.Println(n[12])
	n[12] = "which"
	fmt.Println(n)
	delete(n, 12)
	//通过双赋值检测某个键是否存在：
	//elem, ok = m[key]
	//若 key 在 m 中，ok 为 true ；否则，ok 为 false。
	//若 key 不在映射中，那么 elem 是该映射元素类型的零值。
	//同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
	//注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：
	//elem, ok := m[key]
	fmt.Println(n)
	v, ok := n[4]
	fmt.Println("value is ", v, "存在吗？", ok)

	fmt.Println(m[13])
	fmt.Println(mn["Jayce"])
}
