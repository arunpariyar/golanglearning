// Create your own type. Have the underlying type be an int.
// create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
// in func main
// print out the value of the variable “x”
// print out the type of the variable “x”
// assign 42 to the VARIABLE “x” using the “=” OPERATOR
// print out the value of the variable “x”

package main

import "fmt"

func main() {
	//declaring the type age the underlying type of which is int
	type age int
	var x age = 47
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
