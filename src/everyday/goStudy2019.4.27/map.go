package main

import (
	"fmt"
	"log"
)

//Map

//slcie不能作为map的key
var timeZone = map[string]int{
	"qwe": 12,
}
var timeTwo = map[int]interface{}{
	12: "everyday",
}

func offset(target int) interface{} {
	if value, b := timeTwo[target]; b {
		//b if true
		return value
	}
	//b is false
	log.Println("can't find target on timeTwo")
	return 0

}
func main() {
	//off := timeZone["qwe"]
	//fmt.Print(off)
	fmt.Print(offset(2))
	delete(timeTwo, 13)
}
