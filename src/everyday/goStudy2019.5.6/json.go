package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}
type family struct {
	Name    string
	Age     int
	Parents []string
}

func main() {
	unmarshal()
}
func marshal() {
	//序列化
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"second", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	//Marshal会返回VC的json编码
	/*Marshal函数会递归的处理值。如果一个值实现了Marshaler接口切非nil指针，会调用其MarshalJSON方法来生成json编码。nil指针异常并不是严格必需的，但会模拟与UnmarshalJSON的行为类似的必需的异常。
	否则，Marshal函数使用下面的基于类型的默认编码格式：
	布尔类型编码为json布尔类型。
	浮点数、整数和Number类型的值编码为json数字类型。
	字符串编码为json字符串。角括号"<"和">"会转义为"\u003c"和"\u003e"以避免某些浏览器吧json输出错误理解为HTML。基于同样的原因，"&"转义为"\u0026"。
	数组和切片类型的值编码为json数组，但[]byte编码为base64编码字符串，nil切片编码为null。
	结构体的值编码为json对象。每一个导出字段变成该对象的一个成员*/
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("./filetmp/vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	//Decoder从输入流解码json对象
	//NewDecoder创建一个从r读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	enc := json.NewEncoder(file)
	//decoder是文件句柄和js编码对象之间的桥梁
	//Encode将v的json编码写入输出流，并会写入一个换行符，参见Marshal函数的文档获取细节信息。
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
func unmarshal() {

	data, err := ioutil.ReadFile("./filetmp/card.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	//假定我现在不知道这个json文件的类型格式
	var f interface{}
	_ = json.Unmarshal(data, &f)
	fmt.Println(f) //字段顺序可能发生过改变
	//知道Json的结构
	var m family
	_ = json.Unmarshal(data, &m)
	fmt.Println("Name:", m.Name)
	fmt.Println("Age", m.Age)
	fmt.Println("Parents", m.Parents)
}
