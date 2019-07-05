package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

//json反序列化为结构体
func testStruct2() (ret string, err error) {
	user1 := &User{
		UserName: "user1",
		NickName: "Murphy",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "taidou008@qq.com",
		Phone:    "110",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		err = fmt.Errorf("json.marshal failed, err:", err)
		return
	}

	ret = string(data)
	return
}

func test1() {
	data, err := testStruct2()
	if err != nil {
		fmt.Println("test struct failed, ", err)
		return
	}

	var user1 User
	err = json.Unmarshal([]byte(data), &user1)
	if err != nil {
		fmt.Println("Unmarshal failed, ", err)
		return
	}
	fmt.Println(user1)
}

func main() {
	test1()
	fmt.Println("-----")
}
