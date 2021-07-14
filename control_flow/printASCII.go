package main

import "fmt"

func main() {
	//use a for loop from 33 to 122
	// use format printing to print the ascii values as well
	for i := 65; i <= 90; i++ {
		fmt.Printf("value: %v \t ASCII: %#U \n", i, i)
	}
}