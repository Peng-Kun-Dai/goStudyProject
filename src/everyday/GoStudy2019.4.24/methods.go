package main

//有接收者的函数叫方法
import (
	"fmt"
	"math"
)

/*Go 没有类。不过你可以为结构体类型定义方法。
方法就是一类带特殊的 接收者 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。*/

type Vertex struct {
	X, Y float64
}

//为结构体定义方法
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//为指针接受者声明方法
/*
这意味着对于某类型 T，接收者的类型可以用 *T 的文法。（此外，T 不能是像 *int 这样的指针。）
例如，这里为 *Vertex 定义了 Scale 方法。
指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。
由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
试着移除 Scale 函数声明中的 *，观察此程序的行为如何变化。
若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。（对于函数的其它参数也是如此。）
Scale 方法必须用指针接受者来更改 main 函数中声明的 Vertex 的值。*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * 10
	v.Y = v.Y * 10

}
func methods() {
	v := Vertex{6, 8}
	//fmt.Println(Vertex{3, 4}
	v.Scale(10)
	fmt.Print(v.Abs())
}

/*
你也可以为非结构体类型声明方法。
在此例中，我们看到了一个带 Abs 方法的数值类型 MyFloat。
你只能为在同一包内定义的类型的接收者声明方法，
而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
（译注：就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）
*/
type mFloat float64

//为非结构体类型声明方法
func (f mFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func methodsContinued() {

	f := mFloat(-math.Sqrt2)
	fmt.Println(f.abs())
}

func main() {
	methods()
}
