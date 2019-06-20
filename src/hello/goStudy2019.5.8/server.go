package main

import (
	"fmt"
	"net"
)

//服务端
func main() {
	fmt.Println("start server.....")
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen fail,err", err)
	}
	//监听请求
	for {
		//Accept waits for and returns the next connection to the listener.
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept faild,err:", err)
			//try again
			continue
		}
		//
		go process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		//服务器从连接中获取数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		//s := string(buf[0:n])
		var s string = string(buf[:n])
		//var s [n]byte
		st := []rune(s)
		//var data []rune = []rune(string(buf[0:n]))
		fmt.Println(string(st))

	}
}
