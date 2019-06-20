package main

//slice转json
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func testSlice() {
	var m map[string]interface{}
	var s []map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	s = append(s, m)

	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 29
	m["sex"] = "female"
	m["height"] = "180cm"
	s = append(s, m)

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
	//创建一个文件保存JSON数据
	file, _ := os.OpenFile("./filetmp/sliceToJson.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	//Decoder从输入流解码json对象
	//NewDecoder创建一个从r读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	enc := json.NewEncoder(file)
	//decoder是文件句柄和js编码对象之间的桥梁
	//Encode将v的json编码写入输出流，并会写入一个换行符，参见Marshal函数的文档获取细节信息。
	err = enc.Encode(s)
	if err != nil {
		log.Println("Error in encoding json")
	}
}

func main() {
	testSlice()
	fmt.Println("--------")
}
