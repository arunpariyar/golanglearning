// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

package main

import "fmt"

func main() {
	// callme has been assigned to a function which in turn returns a function that returns a string
	callMe := func() func() string {
		return func() string {
			return "Hello"
		}
	}
	// result now holds the function that returns the string
	result := callMe()
	// calling result the string "Hello " is now displayed to the screeng
	fmt.Println(result())

}
