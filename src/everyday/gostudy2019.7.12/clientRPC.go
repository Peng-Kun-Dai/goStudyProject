package main

//客户端
import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//拨号
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	//调用具体的RPC方法在调用client.Call时，
	// 第一个参数是用点号链接的RPC服务名字和方法名字，
	// 第二和第三个参数分别我们定义RPC方法的两个参数。
	err = client.Call("HelloService.Hello", "data", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
