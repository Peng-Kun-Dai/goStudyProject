package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter:")
	input, err := inputReader.ReadString()

}
