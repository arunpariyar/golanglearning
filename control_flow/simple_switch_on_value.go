package main

import "fmt"

func main() {
	switch 108 {
	case 100, 103, 108:
		fmt.Println("I am 100")
	case 105:
		fmt.Println("I am 105")
	case 107:
		fmt.Println("I am 108")
	default: 
		fmt.Println("There is no 108! Sorry")
	}
}
