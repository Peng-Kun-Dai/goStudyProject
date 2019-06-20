package test

import "fmt"

type shaper interface {
	area() int
}

type big struct {
	a int
}
type small struct {
	b bool
}

func (biger big) area() int {
	return biger.a
}
func (smaller small) area() {
	fmt.Println(smaller.b)
}

func main() {
	//var biger big
	/*//biger := new(big)
	biger := big{5}
	smaller := small{true}
	//var smaller small
	var shap shaper
	shap = &biger
	shap.area()
	shap = &smaller
	shap.area()
	_, ok := shap.(*small)
	fmt.Println(ok)*/

	var shapagain shaper
	var biger big
	shapagain = biger
	switch t := shapagain.(type) {
	case big:
		fmt.Printf("Param #%d is a bool\n", t)
	/*case small:
	fmt.Printf("Param #%d is a float64\n", t)*/
	case nil:
		fmt.Printf("Param #%d is a nil\n", t)
	default:
		fmt.Printf("Param #%d is unknown\n", t)
	}
	var bigchecker interface{} = biger
	if sv, ok := bigchecker.(shaper); ok {
		fmt.Printf("v implements shaper(): \n", sv.area()) // note: sv, not v
	}
	fmt.Println("-------------------")
	(&biger).see()

}

func (biger *big) see() {
	fmt.Println("i am big")
}
