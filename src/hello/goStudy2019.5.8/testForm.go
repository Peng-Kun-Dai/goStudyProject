package main

//表单提交
import (
	"fmt"
	"io"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
                    <input type="text" name="in"/>
                    <input type="text" name="in"/>
                     <input type="submit" value="Submit"/>
             </form></body></html>`

func main() {
	http.HandleFunc("/", fromServer)
	if err := http.ListenAndServe(":8880", nil); err != nil {
		fmt.Println("http listen fail")
	}
}
func fromServer(w http.ResponseWriter, r *http.Request) {
	//头部设置
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.Form["in"][1])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("in"))
	}

}
