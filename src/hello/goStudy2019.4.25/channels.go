package main

/*
信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
（“箭头”就是数据流的方向。）
和映射与切片一样，信道在使用前必须创建：
ch := make(chan int)
默认情况下，发送和接收操作在另一端准备好之前都会阻塞。
这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
*/
import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //将和送入信道
}

func main() {
	s := []int{7, 2, 4, 6, 4, 1}
	//创建信道
	c := make(chan int)
	go sum(s[:len(s)/2], c) //[:3]   {7,2,4}   13
	go sum(s[len(s)/2:], c) //[3:]  {6,4,1}		11
	x, y := <-c, <-c        //可能会阻塞
	fmt.Println("两者和为", x+y)
	fmt.Println("两者差为", x-y)

}
