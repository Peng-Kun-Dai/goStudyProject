package main

//elem反射获取tag
import (
	"fmt"
	"reflect"
)

type Student1 struct {
	Name  string `json:"stu_name"`
	Age   int
	Score float32
}

func TestStruct1(a interface{}) {
	typ := reflect.TypeOf(a)

	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("Tag:%s\n", tag)
}

func main() {
	/*var a = Student1{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
	}*/
	//TestStruct1(&a)
	type User struct {
		Name   string "user name"
		Passwd string `user passsword`
	}
	u := &User{
		Name:   "Murphy",
		Passwd: "123456",
	}
	s := reflect.TypeOf(u).Elem()
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag)
	}
}
