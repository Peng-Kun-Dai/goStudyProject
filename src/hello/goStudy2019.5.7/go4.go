package main

//反射调用结构体方法
import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

func (s student) Print() {
	fmt.Println(s)
}

func (s student) Set(name string, age int, score float32) {
	s.Age = age
	s.Name = name
	s.Score = score
}

func testStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()

	fmt.Println(val, kd)
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取字段数量
	fields := val.NumField()
	fmt.Printf("struct has %d field\n", fields)
	//获取字段的类型
	for i := 0; i < fields; i++ {
		fmt.Printf("%d %v\n", i, val.Field(i).Kind())
	}
	//获取方法数量
	methods := val.NumMethod()
	fmt.Printf("struct has %d methods\n", methods)

	//反射调用的Print方法
	var params []reflect.Value
	val.Method(0).Call(params)

}

func changeStruct(b interface{}) {
	val := reflect.ValueOf(b)
	val.Elem().Field(0).SetString("STU01")
}
func main() {
	var a student = student{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
	}
	//testStruct(a)
	fmt.Println(a)
	changeStruct(&a)
	fmt.Println(a)
}
