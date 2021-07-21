//using closure we are able to enclose the scope of the values in out program where we would like them to be.

package main

import "fmt"

func main() {

	fmt.Println(increment())
	fmt.Println(increment())
}

func increment() int {
	x := 0
	x++
	return x
}
