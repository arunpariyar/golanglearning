package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("you find yourself in a dimly lit cave")
	var command = "walk inside"
	var exit = strings.Contains(command, "outside")
	command = "Emerging, your eyes meet the blinding midday sun"
	var wearShades = strings.Contains(command, "sun")
	fmt.Println("You leave the cave:", exit, "don't forget to wear you shades", wearShades)

	//3.1.2
	command = "the sign at the cave entrance"
	var read = strings.Contains(command, "read")
	fmt.Println("Read the sign", read)
	fmt.Println("apple" < "banana")
}