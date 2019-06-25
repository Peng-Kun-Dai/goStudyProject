package main

import (
	"os"
	"testing"
	"time"
)

//调试学习
//work func
func add(x, y int) bool {
	if x > y {
		return true
	} else {
		return false
	}
}

//Fail FailNow SkipNow
//test func of add
func aTestAdd(t *testing.T) {
	if add(1, 2) != false {
		t.FailNow()
	}
}
func aTest1(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)

}
func aTest2(t *testing.T) {
	if os.Args[len(os.Args)-2] == "2" {
		//if exist 2
		t.Parallel()
	}
	if os.Args[len(os.Args)-1] == "sleep" {
		//if exist 2
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second)
}

func aBenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}
func TestAdd(t *testing.T) {
	if add(1, 2) != false {
		t.Fatal("xxx")
	}
}

type User struct {
	name string
	age  int
}

func (i User) structToMap() map[string]interface{} {
	var data = make(map[string]interface{})
	data["name"] = i.name
	data["age"] = i.age
	return data
}
