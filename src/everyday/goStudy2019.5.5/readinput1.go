package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	input1      string
	err         error
)

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")
	input1, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("the input was :%s\n", input1)
	}
}
