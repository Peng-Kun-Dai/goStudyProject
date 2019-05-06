package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//createfile()
	//readfile()
	//readAllAndWrite()
	//readFileBuffer()
	readFileColumn()
}
func createfile() {
	//文件操作
	//创建文件
	//Create会返回一个文件对象，默认权限0666
	//如果源文件存在将被覆盖
	file1, err1 := os.Create("./filetmp/file1.log")
	if err1 != nil {
		//log.Println(err1)
		log.Fatalln(err1)
	}
	defer file1.Close()
}
func readfile() {
	inputfile, err1 := os.Open("./filetmp/file1.log")
	if err1 != nil {
		log.Fatalln(err1)
	}
	defer inputfile.Close()
	//获取一个读取器变量
	inputReader := bufio.NewReader(inputfile)
	for {
		//逐行的读取
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", (inputString))

		if readerError == io.EOF { /*var p []byte
			inputString, readerError := inputReader.Read(p)*/
			return
		}
	}

}
func readAllAndWrite() {
	//将整个文件的内容读到一个字符串里：
	inputFile := "./filetmp/file1.log"
	outputFile := "./filetmp/file1_copy.log"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x644)
	if err != nil {
		panic(err.Error())
	}
}
func readFileBuffer() {
	//带缓冲的读取
	//buf 的大小即是缓冲的大小
	inputfile, err1 := os.Open("./filetmp/file2.log")
	if err1 != nil {
		log.Fatalln(err1)
	}
	defer inputfile.Close()
	//获取一个读取器变量
	inputReader := bufio.NewReader(inputfile)
	//
	for {
		buf := make([]byte, 1024)
		n, readerError := inputReader.Read(buf)
		fmt.Printf("The input was: %s", buf)

		if readerError == io.EOF && n == 0 { /*var p []byte
			inputString, readerError := inputReader.Read(p)*/
			return

		}
	}

}
func readFileColumn() {
	file3, err := os.Open("./filetmp/file3.txt")
	if err != nil {
		log.Fatalln(err)
		//panic(err)
	}
	defer file3.Close()
	//分别存储三列数据
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		//换行时停止扫描
		_, err = fmt.Fscanln(file3, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
