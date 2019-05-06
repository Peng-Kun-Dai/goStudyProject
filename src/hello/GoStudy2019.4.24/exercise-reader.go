package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

func main() {
	reader.Validate(MyReader{})
}

func (read MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}
