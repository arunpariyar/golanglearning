package main

import "fmt"

func main() {
	slice2D := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Money", "Hello, James"}}
	for index, value := range slice2D {
		fmt.Printf("%v \t %v \n", index, value)
	}
}
