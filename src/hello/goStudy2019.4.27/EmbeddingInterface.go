package main

type eater interface {
	eat()
}

func (f mFloat) eat() {

}

type drinker interface {
	drink()
}

func (f mFloat) drink() {

}

type eatAndDrinker interface {
	eater
	drinker
	fun()
}

func (f mFloat) fun() {

}

type mFloat float64

func main() {
	var f mFloat
	var life eatAndDrinker = f
	_ = life
}
