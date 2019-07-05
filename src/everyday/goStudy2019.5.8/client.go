package main

//客户端
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//拨号
	concn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer concn.Close()
	//input data
	inputReader := bufio.NewReader(os.Stdin)
	for {
		//读取以换行结束
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = concn.Write([]byte(trimmedInput))
		if err != nil {
			return
		}
	}

}
