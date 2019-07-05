package main

import (
	"fmt"
	"sync"
	"time"
)

var s []int
var lock sync.Mutex //互斥锁
func appendValue(i int) {
	lock.Lock() //加锁
	s = append(s, i)
	lock.Unlock() //解锁
}

func main() {
	for i := 0; i < 1000; i++ {
		go appendValue(i)
	}
	//sort.Ints(s) //给切片排序,先排完序再打印,和下面一句效果相同
	time.Sleep(time.Second) //间隔1s再打印,防止一边插入数据一边打印时数据乱序
	for i, v := range s {
		fmt.Println(i, ":", v)
	}
}
