package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	//解析参数，默认不解析
	err := r.ParseForm()
	if err != nil {
		log.Println(r, err)
	}

	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("nal:", strings.Join(v, " "))
	}
	_, _ = fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的

}
func main() {
	http.HandleFunc("/", sayHelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
