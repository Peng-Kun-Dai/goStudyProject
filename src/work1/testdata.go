package main

//常规
//go:generate structToMapTool -type=User,date,test
type User struct {
	name string
	age  int
}

//匿名
type date struct {
	name1, name2 string
	string
}

//嵌套
type test struct {
	user1 *User
	user2 User
	date
}

//空结构体
type empty struct {
}

//组合
//go:generate structToMapTool -type=man,woman
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
