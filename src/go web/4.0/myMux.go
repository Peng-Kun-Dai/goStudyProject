package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloname(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayhelloname(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w.Header(), "hellomyroute")
}
func main() {
	mux := &MyMux{}
	_ = http.ListenAndServe(":9090", mux)
}
