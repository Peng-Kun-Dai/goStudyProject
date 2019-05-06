package main

/*
信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：
ch := make(chan int, 100)
仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞
*/
import (
	"fmt"
	"time"
)

//信道转型，仅发送信道
func write(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		//<-ch
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {
	/*ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)*/
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}
}
