package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//客户端访问百度服务器
func main() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("get error", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data error", err)
		return
	}
	fmt.Println(string(data))
}
