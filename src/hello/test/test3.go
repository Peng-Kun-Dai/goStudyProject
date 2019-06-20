package test

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UTC()
	fmt.Println(t)
	t1 := time.Now().Unix()
	fmt.Println(t1)
	t2 := time.Now().UnixNano()
	fmt.Println(t2)

}
