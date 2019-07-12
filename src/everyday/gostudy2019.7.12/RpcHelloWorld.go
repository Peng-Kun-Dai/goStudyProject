package main

//服务端
import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

//实现打印功能
// go 语言RPC规则：
// 只能有两个可序列化的参数，
// 第二个参数是指针类型，
// 并返回一个error
// 必须是公开的方法
func (p *HelloService) Hello(Request string, reply *string) error {
	*reply = "RPChello:" + Request
	return nil
}

func main() {
	//将HelloService注册为一个RPC服务
	//HelloService对象中所有满足RPC规则的对象方法都会被注册为RPC函数，
	// 所有的注册方法会放在HelloService的服务空间下
	_ = rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("accept error:", err)
	}

	rpc.ServeConn(conn)

}
