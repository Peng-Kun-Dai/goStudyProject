package main

//map转JSON
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*map转json*/

func testMap() {
	var mmp map[string]interface{}
	mmp = make(map[string]interface{})

	mmp["username"] = "Murphy"
	mmp["age"] = 19
	mmp["sex"] = "man"

	data, err := json.Marshal(mmp)
	if err != nil {
		fmt.Println("json marshal failed,err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))
	//创建一个文件保存JSON数据
	file, _ := os.OpenFile("./filetmp/mapToJson.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	//Decoder从输入流解码json对象
	//NewDecoder创建一个从r读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	enc := json.NewEncoder(file)
	//decoder是文件句柄和js编码对象之间的桥梁
	//Encode将v的json编码写入输出流，并会写入一个换行符，参见Marshal函数的文档获取细节信息。
	err = enc.Encode(mmp)
	if err != nil {
		log.Println("Error in encoding json")
	}

}

func main() {
	testMap()
	fmt.Println("----")
}
