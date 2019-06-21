package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(math1(1, 2))
	fmt.Println(mathagain(10, 20))
	fmt.Println(mysqrt(2))
}

//命名返回值
func math1(a, b int) (c, d, e int) {
	c = a + b
	d = a * b
	e = a - b
	return
}
func mysqrt(f float64) (a float64, err error) {
	if f > 0 {
		a = math.Sqrt(f)
	} else {
		a = f
		err = errors.New("i won't be able to do a sqrt of negative")
	}
	return
}

//非命名返回值
func mathagain(a, b int) (int, int, int) {
	c := a + b
	d := a * b
	e := a - b
	return c, d, e
}
