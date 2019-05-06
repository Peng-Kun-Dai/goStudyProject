package main

import (
	"fmt"
	"time"
)

//错误信息的结构
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	//return fmt.Sprintf("at %v, %s", e.When, e.What)
	return fmt.Sprint(*e)
}

//因为这里返回了error类型的数据，所以必须重写error类的方法Error()
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
