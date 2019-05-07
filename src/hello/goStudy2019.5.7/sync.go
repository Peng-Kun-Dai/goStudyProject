package main

import (
	"fmt"
	"sync"
)

type info struct {
	sync.Mutex
	age int
	m   map[int]string
}

func main() {
	var in info
	//runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go f(&in, i)
	}
	wg.Wait()

}
func f(in *info, i int) {
	in.Lock()
	in.age = i
	in.Unlock()
	fmt.Println(i)
}
