// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

var favSport string

func main() {
	favSport = "lala"

	switch favSport {
	case "football":
		fmt.Println("Are you a messi fan ?")
	case "badminton":
		fmt.Println("Is that even a game")
	case "skateboard":
		fmt.Println("My man !!")
	default:
		fmt.Println("Sorry your sport is not familiar to us")
	}

}
