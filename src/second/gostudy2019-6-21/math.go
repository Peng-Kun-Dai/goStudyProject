package main

import "runtime"

func main() {
	println(runtime.NumCPU())

	go func() {
		func() {

		}()
	}()
}