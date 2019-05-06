package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type book struct {
	titlt    string
	price    float64
	quantity int
}

func main() {

	books := make([]book, 1)
	file, err := os.Open("./filetmp/products.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// remove \r and \n so 2
		line = string(line[:len(line)-2])

		//用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，
		// 即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
		str1 := strings.Split(line, ";")
		book := new(book)
		book.titlt = str1[0]
		book.price, err = strconv.ParseFloat(str1[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		//Atoi是ParseInt(s, 10, 0)的简写。
		book.quantity, err = strconv.Atoi(str1[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		//将一组数据放入切片
		if books[0].titlt == "" {
			books[0] = *book
		} else {
			books = append(books, *book)
		}

	}
	fmt.Println("切片内容：")
	for _, bk := range books {
		fmt.Println(bk)
	}

}
