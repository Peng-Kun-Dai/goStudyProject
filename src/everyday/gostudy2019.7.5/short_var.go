package main

//使用短声明导致变量覆盖
import "fmt"

func main() {
	/*	var remember bool = false
		if true {
			remember := true //两个remember不一致
		}*/
	//use remember
}
func shadow() (err error) { //err已经被声明
	x, err := check1() // x是新创建变量，err是被赋值
	if err != nil {
		return // 正确返回err
	}
	if y, err := check2(x); err != nil { // y和if语句中err被创建
		return // if语句中的err覆盖外面的err，所以错误的返回nil！
	} else {
		fmt.Println(y)
	}
	return
}
