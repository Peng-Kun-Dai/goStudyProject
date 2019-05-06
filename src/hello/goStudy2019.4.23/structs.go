package main

import "fmt"

type people struct {
	name   string
	sex    bool
	height int
}

func main() {
	//struct1()
	struct3()
}

//结构体初始化以及修改访问
func struct1() {
	fmt.Println(people{"代鹏坤", true, 180})
	i := people{"代鹏坤", true, 180}
	//rename
	i.name = "小代"
	fmt.Println(i)
	fmt.Println(i.name)
}

/*如果我们有一个指向结构体的指针 p，
那么可以通过 (*p).X 来访问其字段 X。
不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，
直接写 p.X 就可以。
*/
func struct2() {
	i := people{"代鹏坤", true, 180}
	p := &i
	//change
	(*p).height = 190
	p.name = "Jayce"
	fmt.Println(*p)
}
func struct3() {

	var (
		i1 = people{"代鹏坤", true, 180}
		//Name:语法可以仅列出部分字段。（字段名的顺序无关。）
		i2 = people{name: "jayce", height: 180} //性别被隐式的赋值为false
		i3 = people{}                           //所有属性全部隐式的赋予
		p  = &people{"代鹏坤", true, 180}
	)

	fmt.Println(i1, p, i2, i3)

}
