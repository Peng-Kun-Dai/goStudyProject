package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {

}
func gob1() {
	var network bytes.Buffer

	enc := gob.NewEncoder(&network) // Will write to network.
	err := enc.Encode(P{3, 4, 5, "Jayce"})
	//Encode方法将e编码后发送，并且会保证所有的类型信息都先发送。
	if err != nil {
		log.Fatal(err)
	}

	//Decode从输入流读取下一个之并将该值存入e。如果e是nil，将丢弃该值；否则e必须是可接收该值的类型的指针。
	// 如果输入结束，方法会返回io.EOF并且不修改e（指向的值）。
	dec := gob.NewDecoder(&network) // Will read from network.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%q:{%d,%d}\n", q.Name, *q.X, *q.Y)
}
func gob2() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
