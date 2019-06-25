package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) //计数器加一
		go func(i int) {
			defer wg.Done() //计数器减一
			time.Sleep(time.Millisecond)
			println(i)
		}(i)
	}
	wg.Wait()
	println("main exit..")
}
