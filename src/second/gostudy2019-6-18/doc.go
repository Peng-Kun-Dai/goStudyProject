/*包注释*/
package main

import (
	"log"
	"os"
)

/*问题*/
/*cwd全局变量未被初始化*/
var cwd string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.getwd failed: %v", err)
	}
	log.Printf("working directory = %s", cwd)
}

/*解决办法*/
var cwd1 string

func init() {
	var err error
	cwd1, err = os.Getwd()
	if err != nil {
		log.Fatal("os.getwd failed: %v", err)
	}
}
