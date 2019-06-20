package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputerror := os.Open("input.dat")
	if inputerror != nil {
		fmt.Println("cann't find file")
		return
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Println("the input was: %s", inputString)
	}
}
