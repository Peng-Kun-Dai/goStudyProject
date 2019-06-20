package main

//定时器是进程规划⾃⼰在未来某⼀时刻接获通知的⼀种机制。
import (
	"fmt"
	"time"
)

func main() {
	//timer()
	after()
}

//方法一  Ticker间隔特定时间触发
func ticker() {
	t := time.NewTicker(time.Second * 3) //3三秒执行一次
	// t := time.NewTicker(执行周期)
	for v := range t.C {
		fmt.Println("hello, ", v)
	}
}

//方法2  Timer （到达指定时间触发且只触发⼀次）
func timer() {
	start := time.Now()
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("after func callback, elaspe:", time.Now().Sub(start))
	})

	time.Sleep(1 * time.Second)
	//time.Sleep(3 * time.Second)
	//timer.Stop()
	// Reset 在 Timer 还未触发时返回 true；触发了或Stop了，返回false
	//将timer激活并且等待3秒后执行
	if timer.Reset(3 * time.Second) {
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop!")
	}

	time.Sleep(10 * time.Second)
}

//模拟超时
func after() {
	start := time.Now()
	select {
	case <-time.After(time.Second * 5): // 5秒后执行
		// case <- time.After(周期):
		fmt.Println("5s after")
	case <-time.After(time.Second * 3): // 3秒后执行
		// case <- time.After(周期):
		fmt.Println("3s after")
	}
	duration := time.Since(start)
	fmt.Println("运行时间：", duration)
}
