package main

//运行时间
import (
	"fmt"
	"time"
)

func main() {
	startCac()
}

//计算程序花费时间
func startCac() {
	t1 := time.Now()
	time.Sleep(2 * time.Second)
	time1 := time.Since(t1)
	time.Sleep(2 * time.Second)
	time2 := time.Since(t1)
	fmt.Println(time1, "\n", time2)
}
