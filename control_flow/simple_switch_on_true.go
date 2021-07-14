package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("I am false")
	case false:
		fmt.Println("I am false")
	case false:
		fmt.Println("I am true")
	default:
		fmt.Println("I am the default case")
	}
}
