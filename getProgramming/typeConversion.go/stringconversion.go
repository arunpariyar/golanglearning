package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 40

	fmt.Println("I am " + strconv.Itoa(age) + " years old")
	doB := "1990"
	dateOfBirth , err := strconv.Atoi(doB)
	if err != nil{ 
		fmt.Println("error", err)
	}
	fmt.Println("I was born in ", dateOfBirth)
}
