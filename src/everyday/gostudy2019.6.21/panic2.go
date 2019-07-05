package main

import "log"

func main() {
	defer func() {
		log.Fatal(recover())
	}()
	panic2()
}
func panic2() {
	defer panic("i am dead")
	panic("i am dead again")
}
