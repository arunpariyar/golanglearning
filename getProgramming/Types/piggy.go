/* Write a program that randomly places nickels ($0.05), dimes ($0.10), and quarters ($0.25) into an empty piggy bank until it con- tains at least $20.00. Display the running balance of the piggy bank after each deposit, formatting it with an appropriate width and precision. */
package main

import (
	"fmt"
	"math/rand"
)

func main(){
	//intialize piggybank
	var piggybank float64 = 0.0
	//until its adds to 20
	for piggybank < 20.00 {
		switch rand.Intn(3)+1 {
		case 1:
			piggybank += 0.05
		case 2:
			piggybank += 0.10
		case 3: 
			piggybank += 0.25
		}
		fmt.Printf("%5.2f \n", piggybank)
	}
}