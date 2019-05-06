package main

/*
接口也是值。它们可以像其它值一样传递。
接口值可以用作函数的参数或返回值。
在内部，接口值可以看做包含值和具体类型的元组：
(value, type)
接口值保存了一个具体底层类型的具体值。
接口值调用方法时会执行其底层类型的同名方法。
*/
import (
	"fmt"
	"math"
)

//接口s声明
type J interface {
	M()
	//N()
}

//结构体类型以及M的实现
type TT struct {
	S string
}

func (t *TT) M() {
	if t == nil {
		fmt.Println("t为<nil>，but j 不为nil,包含一个空t")
		return
	}
	fmt.Println(t.S)
}

//其他类型以及对M的实现
type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(j J) {
	fmt.Printf("(%v, %T)\n", j, j) //值，类型
}
func main() {
	var j J

	j = &TT{"Hello"} //对应类型必须实现接口中的全部方法
	describe(j)
	j.M()
	//s:=j.(TT)

	j = F(math.Pi)
	describe(j)
	j.M()
	s := j.(F)
	fmt.Println(s)
}
