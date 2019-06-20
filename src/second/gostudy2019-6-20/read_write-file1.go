package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "./filetmp/products.txt"
	outputFile := "./filetmp/products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile) //file目录
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		//panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x644)
	if err != nil {
		panic(err.Error())
	}
}
