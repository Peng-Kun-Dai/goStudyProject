package main

//常规
//go:generate structToMap -type=User
type User struct {
	name string
	age  int
}

//匿名
//go:generate structToMap -type=date
type date struct {
	name1, name2 string
	string
}

//嵌套
//go:generate structToMap -type=test
type test struct {
	user1 *User
	user2 User
	date
}

//组合
//go:generate structToMap -type=man
//go:generate structToMap -type=woman
type (
	man struct {
		name string
		age  int
	}
	woman struct {
		name string
		age  int
	}
)

func main() {
}

type mInt int
type ha struct {
}
