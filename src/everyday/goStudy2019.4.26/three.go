package main

import "fmt"

func main() {
	//var oldMap  = make(map[int]string)
	var oldMap map[int]string
	oldMap = make(map[int]string)
	oldMap[1] = "everyday"
	fmt.Println(oldMap[1])

	//测试copymap
	fmt.Println("测试copymap")
	var newmap map[int]string = copymap(oldMap)
	fmt.Println(newmap[1])

	//测试copykeymap
	fmt.Println("测试copykeymap")
	key := copykeymap(oldMap)
	fmt.Println(key)
	fmt.Println(oldMap[1])

	//测试testRange
	fmt.Println("初始值为", oldMap)
	testRange(oldMap)
	fmt.Println("测试后的值为", oldMap)

}

func copymap(oldmap map[int]string) map[int]string {
	var newmap map[int]string = make(map[int]string)
	for key, value := range oldmap {
		newmap[key] = value
	}
	return newmap
}

func copykeymap(oldmap map[int]string) int {
	//取第一项的值
	for key := range oldmap {
		oldmap[1] = "world"
		//delete(oldmap, key)
		return key
	}
	//oldmap为空的情况
	return 0
}
func copyvaluemap(oldmap map[int]string) string {
	//取第一项的值
	for _, value := range oldmap {
		//oldmap[1] = "world"
		//delete(oldmap, key)
		return value
	}
	//oldmap为空的情况
	return ""
}

func testRange(oldmap map[int]string) {
	for _, value := range oldmap {
		value = "new everyday"
		fmt.Println(value)
	}
}
