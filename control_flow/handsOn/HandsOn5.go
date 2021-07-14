package main

import "fmt"

func main() {
	// Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

	//for loop to go throught 10 and 100
	for i := 10; i <= 100; i++ {
		//use module to divide 4
		modulo := i % 4
		//print the remainder
		fmt.Printf("%v \n", modulo)

	}

}
