package main

import (
	"fmt"
	"time"
)

//并发
//拒绝以共享实现通信  难以保证共享变量的正确性  修改时加锁
//提倡以通信实现共享	channel

//goroutines
//多个goroutines之间资源共享，并且非常轻量
//goroutines 与操作系统线程  “多对多”，一个被阻塞时，其他依然能够运行
//函数执行结束，goroutine隐式退出
//函数中被引用的变量在函数结束以前不会被释放
func announce(message string, delay time.Duration) {
	//闭包，这个函数是没有名字的
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}()
}

//channel
//make分配，类型，缓冲区大小
func testchannel() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
			c <- 1
		}
	}()
	fmt.Println("退出testchannel")
	<-c
}

//循环的迭代变量会在所有goroutine中共享，这不是我们乐意的
//要保证每个request是每个goroutine私有的
//req:=req
//
//
//
type request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

//------------------------------
type vector []float64

/*func (v vector) doSome(i, n int, u vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
}*/
func main() {
	testchannel()
	//v := vector{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}
	//c := make(chan int)
	//fmt.Println(v.doSome(2, 3, v, c))
	//panic("再见")
}
