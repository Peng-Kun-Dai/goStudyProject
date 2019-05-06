package main

import "fmt"

func main() {
	var se sequence = []int{1, 2, 3}
	//如果使用指针接受者来实现一个方法，那么只有指向类型的指针才能实现对应的方法。
	// 如果使用值接收者来实现一个方法，那么这个类型的值和指针都能够实现对应的方法。
	var inter interfacename = &se
	//var inter2  interfacename= se
	//var object,object2 inferfaceObject  =se,&se

	inter.swap(0, 1)
	fmt.Print(inter.less(0, 1))
}

type sequence []int

//空接口
type inferfaceObject interface {
}

//这个类型的接口只接受实现了实现了声明方法的类型的值
type interfacename interface {
	len() int
	less(i, j int) bool
	swap(i, j int)
}

//指针接收者
func (s *sequence) len() int {
	return len(*s)
}

//值接收者
func (s sequence) less(i, j int) bool {
	return s[i] < s[j]
}
func (s sequence) swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//ifunc (s sequence) String() string {
//	sort.IntSlice(s).Sort()
//	str := "["
//	for i, elem := range s {
//		if i > 0 {
//			str += " "
//		}
//		str += fmt.Sprint(elem)
//	}
//	return str + "]"
//}
