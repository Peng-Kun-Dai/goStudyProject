package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	//mkdir_json()
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc) //js is a []byte
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("./json/vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}

func mkdir_json() {
	err2 := os.MkdirAll("./json", os.ModePerm)
	if err2 != nil {
		//Fatal等价于{l.Print(v...); os.Exit(1)}
		//log.Fatal(err)
		log.Println(err2)
	}
}
