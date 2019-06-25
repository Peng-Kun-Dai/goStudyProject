package test

//常规
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

//组合
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
