package main

//elem反射操作结构体
import (
	"fmt"
	"reflect"
)

type sStudent struct {
	Name  string
	Age   int
	Score float32
}

func (s sStudent) Print() {
	fmt.Println(s)
}

func TestStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()

	fmt.Println(val, kd)
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取字段数量
	fields := val.Elem().NumField()
	fmt.Printf("struct has %d field\n", fields)
	//获取字段的类型
	for i := 0; i < fields; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}
	//获取方法数量
	methods := val.NumMethod()
	fmt.Printf("struct has %d methods\n", methods)

	//反射调用的Print方法
	var params []reflect.Value
	val.Elem().Method(0).Call(params)
}

func main() {
	var a sStudent = sStudent{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
	}
	TestStruct(&a)
	// fmt.Println(a)
}
