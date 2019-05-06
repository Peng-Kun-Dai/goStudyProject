package main

import "fmt"

type transform [3][3]float64
type lineOfText [][]byte

//分配二维数组
/*func f1() {
	text := lineOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party"),
	}
}
func f2() {
	picture := make([][]uint8, Ysize)
	for i := range picture {
		picture[i] = make([]uint8, Xsize)
	}
}
func f3() {
	picture := make([][]uint8, Ysize)

	pixels := make([]uint8, Xsize*Ysize)
	for i := range picture {
		picture[i], pixels = pixels[:Xsize], pixels[Xsize:]
	}
}
*/

func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	//x = append(x, 4, 5, 6)
	x = append(x, y...) //x=append(x,y)编译会报错
	fmt.Print(x)

}
