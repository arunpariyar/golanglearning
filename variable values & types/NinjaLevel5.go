package main

import "fmt"

//declaring the type age the underlying type of which is int
	type age int
	var x age = 47.78
	var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y = int(x)
	
	fmt.Println(y)
	fmt.Printf("%T \n", y)
}