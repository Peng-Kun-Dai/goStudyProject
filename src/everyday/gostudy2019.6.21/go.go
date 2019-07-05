package main

import "time"

func add() int {
	return 10
}
func main() {
	go add()

	go func(x, y int) {
		time.Sleep(time.Millisecond)
		println("go:", x, y)
	}(10, add())
	_ = make(map[string]int)

}
