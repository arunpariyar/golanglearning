package main

import "fmt"

func main() {
	if true {
		fmt.Println("true") // this will run
	}
	if false { 
		fmt.Println("false") // this won't run
	}
	if !true {
		fmt.Println("!true") // this won't run
	}
	if !false { 
		fmt.Println("!false") // this will run
	}
}
