package main

import "log"

func test() {
	defer println("test.1")
	defer println("test.2")
	panic("i am dead")
}
func main() {
	defer func() {
		log.Println(recover())
	}()
	test()
}

func testdeferpanic() {
	defer func() {
		if error2 := recover(); error2 != nil {
			log.Fatal(error2)
		}
	}()
	panic("i an dead")
	print("nxlczjn")
}
