package main

/*
=号就是拷贝
1. golang 切片和struct的赋值为值拷贝，map为引用拷贝。
2. 如果要做到修改对象，使用数组指针。
*/
import "fmt"

//slice1:=[]int{1,2,3}

//slice的拷贝问题
func slice() {
	//make map
	m := make(map[string]interface{}, 0)
	//make slice
	s := make([]int64, 1, 4)
	//给map中添加一组值
	m["hello"] = s
	//slice append
	s = append(s, 1) //append之后的slice已经不是之前的slice了
	//证明append会创建一个新的slice
	//看看map中的slice有没有发生变化
	fmt.Print(m["hello"])

	fmt.Println("测试2")
	a := make([]int, 0, 10)
	a2 := a //此时a2,a指向同一片底层数组
	fmt.Printf("%p\n", &a)
	//fmt.Println(a.)
	fmt.Printf("%p\n", &a2)
	a = append(a, 1)
	a2 = append(a2, 2)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &a2)
	//a2 = append(a2, 20)
	//a2[0] = 20
	fmt.Println(a, a2)

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t := data[0:2:3] //长度为2，容量为3
	fmt.Printf("len: %d,cap: %d\n", len(s), cap(s))
	fmt.Printf("原地址对比\n%p\n%p\n", &data[0], &t[0])
	t2 := append(t, 100) // 添加1一个值，未超过容量
	s[0] = 100
	fmt.Printf("len: %d,cap: %d\n", len(t2), cap(t2))
	fmt.Printf("未超过容量时地址对比\n%p\n%p\n", &data[0], &t2[0])
	t3 := append(t2, 200) // 再次添加一个值，未超过容量3
	fmt.Printf("len: %d,cap: %d\n", len(t3), cap(t3))
	fmt.Printf("超过容量时地址对比\n%p\n%p\n", &data[0], &t3[0])
}

//map的值传递
func mapCopy() {
	m := make(map[string]interface{}, 0)
	ll := map[string]interface{}{
		"hello": "hi",
	}
	m["hello"] = ll //因为赋值时传递是一个指针  引用拷贝
	ll["hi"] = 1
	fmt.Println(m["hello"]) //map[hello:hi hi:1]

	m2 := ll
	m2["hi"] = 2
	fmt.Println(ll, m2)
}

type a struct {
	B int
}

func structCopy() {
	m := make(map[string]interface{}, 0)
	ll := a{1}
	m["hello"] = ll //发生了值拷贝
	ll.B = 10
	fmt.Println(m["hello"])

	x := a{2}
	x2 := x //发生了值拷贝
	x.B = 20
	fmt.Println(x, x2)
}

func main() {
	slice()
	//structCopy()
}
