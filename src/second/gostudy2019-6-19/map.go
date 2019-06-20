package main

import "fmt"

func main() {
	map1 := make(map[int]string)
	map1[3] = "jayce"
	map1[2] = "wene"
	delete(map1, 2)             //删除不存在的key不会报错
	if val, ok := map1[2]; ok { //是否存在
		fmt.Println(val)
	} else {
		fmt.Println("不存在")
	}

	/*for key, value := range map1 { //遍历

	}*/
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}
