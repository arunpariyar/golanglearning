package main

import "fmt"

func main() {
	//use for loop to go through 1 to 100
	for i := 1; i <= 10; i++ {
		//if value % 2 = 0 print even
		if i%2 == 0 {
			fmt.Printf("%v is even \n", i)
			// else print odd
		} else {
			fmt.Printf("%v is odd \n", i)
		}
	}
}
