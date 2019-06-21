package main

import "fmt"

//实现信号量
type empty interface{}
type semaphore chan empty

func main() {
	//可用资源数N
	//sem := make(semaphore, N)
	var ch = make(chan int, 4)
	//close(ch)
	ch <- 1
	fmt.Println(cap(ch))
	fmt.Println(len(ch))
	fmt.Println(<-ch)

}

// acquire n resources
func (s semaphore) P(n int) {
	e := new(empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resouces
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
