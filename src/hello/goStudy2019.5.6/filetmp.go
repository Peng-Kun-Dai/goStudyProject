package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	//在项目目录下创建
	//Mkdir使用指定的权限和名称创建一个目录。如果出错，会返回*PathError底层类型的错误。
	//mkdir()
	fmt.Println("-----第二次创建------")

	//MkdirAll使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误。
	// 权限位perm会应用在每一个被本函数创建的目录上。如果path指定了一个已经存在的目录，MkdirAll不做任何操作并返回nil。
	//如果目录已存在也不会报错
	mkdirall()

	//删除目录
	//如果目录下有文件或其他目录会出错
	remove()

	//删除多级目录
	//如果是单个名称，则删除所有的子目录
	//removeall()
	//等个3秒，看目录是否创建成功
	time.Sleep(time.Second * 3)
}
func mkdir() {
	//在项目目录下创建
	//Mkdir使用指定的权限和名称创建一个目录。如果出错，会返回*PathError底层类型的错误。
	err := os.Mkdir("./filetmp/tmpb", os.ModePerm)
	if err != nil {
		//Fatal等价于{l.Print(v...); os.Exit(1)}
		//log.Fatal(err)
		log.Println(err)
	}
}
func mkdirall() {
	//MkdirAll使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误。
	// 权限位perm会应用在每一个被本函数创建的目录上。如果path指定了一个已经存在的目录，MkdirAll不做任何操作并返回nil。
	//如果目录已存在也不会报错
	err2 := os.MkdirAll("./filetmp/tmpa", os.ModePerm)
	if err2 != nil {
		//Fatal等价于{l.Print(v...); os.Exit(1)}
		//log.Fatal(err)
		log.Println(err2)
	}

}
func remove() {
	//删除目录
	//如果目录下有文件或其他目录会出错
	err3 := os.Remove("./filetmp")
	if err3 != nil {
		log.Println(err3)
	}
}
func removeall() {
	//删除多级目录
	//如果是单个名称，则删除所有的子目录
	err4 := os.RemoveAll("./filetmp")
	if err4 != nil {
		log.Println(err4)
	}
}
