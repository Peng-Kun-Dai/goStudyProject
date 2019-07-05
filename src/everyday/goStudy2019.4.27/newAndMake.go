package main

//new struct。。不初始化内存，只是将其置零，返回它的地址
//make适用于slice，map，channel，返回一个已经初始化的，类型为T的值
import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

type syncedbuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
	mint   //类似继承
}

type mint struct {
	a interface{}
	b int
}

func main() {
	//var test1 syncedbuffer
	//test := new(syncedbuffer)
	//fmt.Println(test1)
	//fmt.Println(test)

	var p []int = make([]int, 10)
	var v *[]int = new([]int)
	fmt.Println(p)
	fmt.Println(*v)
	p = make([]int, 9)
	*v = make([]int, 5)
	fmt.Println(p)
	fmt.Println(*v)

}

func newFile(fd int, name string) *os.File {
	if fd < 0 {
		return nil
	}
	f := new(os.File)
	return f
}
