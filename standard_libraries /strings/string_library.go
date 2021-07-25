package main

import (
	"fmt"
	"strings"
)

func main(){
	//to check if a string contains a substring
	fmt.Println(strings.Contains("test","es"))

	//to count the number of times a smaller string occurs in a big strings
	fmt.Println(strings.Count("hello", "l"))

	//to compare two strings
	fmt.Println(strings.Compare("a","z"))
	
	//Report with yes or no if the string contains any of the Unicode points in chars
	fmt.Println(strings.ContainsAny("James Bond","p"))

	//split the given string to slices based on given char
	brokenDown := strings.Split("a one a two a one two three fours"," ")

	for _, value := range brokenDown { 
		fmt.Println(value)
	}

	fmt.Printf("%T\n", brokenDown)

	name := "   jAmeS Bond  "
	fmt.Println(name)
	name = strings.Trim(name, " ")
	fmt.Println(name)
	
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, "james", "jade")
	fmt.Println(name)


}