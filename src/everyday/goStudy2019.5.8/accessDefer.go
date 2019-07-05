package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://taobao.com",
}

func main() {
	for _, v := range url {
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (conn net.Conn, e error) {
					//延迟两秒进行下一次访问
					timeout := time.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed,err:%v\n", v, err)
			continue
		}
		fmt.Printf("head %s succ, status:%v\n", v, resp.Status)

	}
}
