package main

import (
	"fmt"
	"sync"
	"time"
)

var locks sync.Mutex

func main() {
	m := make(map[int]int)
	go func() { //开一个协程写map
		for i := 0; i < 100; i++ {
			locks.Lock() //加锁
			m[i] = i
			locks.Unlock() //解锁
		}
	}()
	go func() { //开一个协程读map
		for i := 0; i < 100; i++ {
			locks.Lock() //加锁
			fmt.Println(m[i])
			locks.Unlock() //解锁
		}
	}()
	time.Sleep(time.Second * 20)
}
