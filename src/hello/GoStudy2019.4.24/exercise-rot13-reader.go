package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 110:
			b[i] -= 13
		case b[i] >= 97:
			b[i] += 13
		case b[i] >= 78:
			b[i] -= 13
		case b[i] >= 65:
			b[i] += 13

		}
		//b[i] = rot13(b[i])
	}
	//fmt.Print(b)
	return n, err
}

//字母转换
func rot13(out byte) byte {
	switch {
	case out >= 'A' && out <= 'M' || out >= 'a' && out <= 'm':
		out += 13
	case out >= 'N' && out <= 'Z' || out >= 'n' && out <= 'z':
		out -= 13
	}
	return out
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
