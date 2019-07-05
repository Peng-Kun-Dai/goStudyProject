package main

import (
	"fmt"
	"reflect"
)

type people struct {
	name   string
	age    int
	height float64
}

var girl *people = new(people)

func main() {
	//create a struct variable
	boy := new(people)
	boy.name = "jayce"
	boy.age = 12
	boy.height = 40.2
	fmt.Println(*boy)
	//luci := people{"luci", 12, 12.3}
	tree()
}

//递归结构体

type node struct {
	rightnode *node   //左节点
	leftnode  *node   //右节点
	date      float64 "an tag"
	int
	float64
}

func tree() {

	left := node{date: 0.5}
	right := node{date: 1.5}
	root := node{date: 1.0, leftnode: &left, rightnode: &right}
	fmt.Println(root.leftnode)
	i := reflect.TypeOf(root)
	fiel := i.Field(2)
	fmt.Println(fiel.Tag)

}
