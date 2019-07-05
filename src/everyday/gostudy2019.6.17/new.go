package main

import "fmt"

/*new函数创建变量
  创建匿名变量
  返回变量地址*/

func main() {
	p := new(int)   //p,*int 类型，指向匿名的interesting变量
	fmt.Println(*p) //"0"
	*p = 2          //设置匿名变量的值为2
	fmt.Println(*p)
	fmt.Println(new(int) == new(int))
}
