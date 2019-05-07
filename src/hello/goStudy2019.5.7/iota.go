package main

import (
	"fmt"
)

const (
	A, B = iota, iota << 10 //0,0<<10
	C, D                    // 1, 1 << 10
)

func main() {
	fmt.Println(A, B, C, D)
}
