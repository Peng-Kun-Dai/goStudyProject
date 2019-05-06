package main

//switch用法
import "fmt"

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func testswitch() {
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i < 5:
			fmt.Println(i, "小于5哦")
		case i >= 5:
			fmt.Println(i, "大于等于5哦")
			//goto loop   调到loop的位置继续执行代码
			break loop //跳出for

		}

	}
}

func compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
		switch {
		case len(a) > len(b):
			return 1
		case len(b) > len(a):
			return -1
		}
		return 0
	}
	return 0
}
func main() {
	//fmt.Print(unhex('Z'))
	testswitch()
}
