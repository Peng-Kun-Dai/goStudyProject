package main

//浏览器访问服务器
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("localhost:8880", nil)
	if err != nil {
		fmt.Println("httplisten fail")
	}
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("终端输出：handle everyday")
	fmt.Fprint(w, "-------everyday")
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("终端输出：handle login")
	fmt.Fprint(w, "-------login")
}
