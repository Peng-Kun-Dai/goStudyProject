package main

import (
	"errors"
	"log"
)

var errorfound error = errors.New("this is my error,sorry")

func main() {
	err := errorfound
	if err != nil {
		log.Fatal(err)
	}
}
func protect(g func()) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	g() //   possible runtime-error
}
