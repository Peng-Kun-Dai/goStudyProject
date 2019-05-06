package main

import "fmt"

type mInt struct {
	name string
}
type interone interface {
	see()
}

func (m mInt) see() {
	fmt.Print("实现了look")
}

type intertwo interface {
	help()
}

func (m mInt) help() {

}

type node struct {
	data interface {
		see()
	}
}

func main() {

	/*//这里的t是一个匿名的接口类型
	var t interface {
		see()
	} = mInt{"hello"}
	n := node{t}
	n.data.see()*/

	//现将这个类型放入一个空接口中载判断，
	/*var intercheck = mInt{"hello"}
	if v, ok := intercheck.(interone); ok {
		fmt.Printf("%v 实现了 intertwo接口", v)
	}*/
	//会出现以下错误
	//Invalid type assertion: intercheck.(intertwo) (non-interface type mInt on left)
	//确定mInt这个类型是否实现了intertwo
	//方法一
	var intercheck interface{} = mInt{"hello"}
	if v, ok := intercheck.(interone); ok {
		fmt.Printf("%v 实现了 interone接口", v)
	}
}

//方法二
var _ interone = (mInt)(nil)

func test() {
	fmt.Println(text(5.0))
	//知道类型的情况下
	var i interface{} = "hello"

	//comma,ok  安全测试
	s, ok := i.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Print("value is not string")
	}

}
func text(m interface{}) interface{} {
	//m := 5.0
	var interfacevalue interface{} = m
	//类型断言
	switch value := interfacevalue.(type) {
	case int:
		return value
	default:
		return "不确定的类型"

	}
}
