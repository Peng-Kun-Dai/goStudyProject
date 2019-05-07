package main

//Elem反射操作基本类型
import (
	"fmt"
	"reflect"
)

func main() {

	var b int = 1
	b = 200
	testInt(&b)
	fmt.Println(b)
}

//fv.Elem()用来获取指针指向的变量
func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
	// 如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
	val.Elem().SetInt(100)
	//fmt.Println(val.Elem().CanSet())
	c := val.Elem().Int()

	fmt.Printf("get value  interface{} %d\n", c)
	fmt.Printf("string val:%d\n", val.Elem().Int())
}
