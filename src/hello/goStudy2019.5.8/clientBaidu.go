package main

//tcp 获取网站数据
import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("error dialing", err.Error())
		return
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\r\n"        //16
	msg += "Host:www.baidu.com\r\n"    //20
	msg += "Connection:keep-alive\r\n" //23
	msg += "\r\n\r\n"                  //4

	//io.WriteString(os.Stdout, msg)
	//func WriteString(w Writer, s string) (n int, err error)
	//WriteString函数将字符串s的内容写入w中。如果w已经实现了WriteString方法，函数会直接调用该方法。
	n, err := io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed, ", err)
		return
	}
	fmt.Println("send to baidu.com bytes:", n)
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		fmt.Println("count:", count, "err:", err)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))
	}
}
