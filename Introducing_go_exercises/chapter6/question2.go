// Write a function that takes an integer and halves it and returns true if it was even or false if it was odd. For example, half(1) should return (0, false) and half(2) should return (1, true).

package main

import "fmt"

func main() {
	modulo, result := half(4)
	fmt.Println(modulo, result)
}

func half(value int) (int, bool) {
	return value/2 , value %2 ==0  
}
