package main

import (
	. "fmt"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	Println(r.Form) //打印在服务器端
	println("path", r.URL.Path)
	println("scheme", r.URL.Scheme)
	println(r.Form["url_long"])
}
