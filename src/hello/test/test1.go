package test

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 0 || i == 9 {
			for j := 0; j < 20; j++ {
				fmt.Print("* ")
			}
			fmt.Println()
		} else {
			fmt.Print("* ")
			for j := 0; j < 18; j++ {
				fmt.Print("  ")
			}
			fmt.Println("*")
		}

	}
}
