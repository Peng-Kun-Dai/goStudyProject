package main

import (
	"fmt"
	"time"
)

func main() {
	strs := []string{"one", "two", "three"}
	for _, s := range strs { //s是三个线程间共享的变量，最后导致输出结果全部一样
		//s := s

		go func() {

			time.Sleep(1 * time.Second)

			fmt.Printf("%s ", s)

		}()

	}

	time.Sleep(3 * time.Second)
}
