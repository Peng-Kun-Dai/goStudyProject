package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(1000)
	end := time.Now()
	timer := end.Sub(start)
	fmt.Println(timer)
}
