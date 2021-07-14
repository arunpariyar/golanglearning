//Create a program that uses a switch statement with no switch expression specified.

package main

import "fmt"

func main(){

	a := 30 
	b := 30

	switch { 
	case a > b: 
	fmt.Println("a is greater than b")
	case a < b: 
	fmt.Println("a is less than b")	
	default : 
	fmt.Println("a and b are equal ")
	}
}