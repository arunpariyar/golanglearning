//Variable Decleration and infered Typing

package main

import "fmt"

var(
	//Variable declaration using implicit type
	dateOfBirth int

	//grouping together
	firstname, lastname string

	//variable declaration with initialiaers
	address string = "53 Elmers street "
)
	//variable declaration with inferred typing using single var statement

	var houseNo = 54 

func main(){
	firstname = "James"
	lastname = "Wild"
	dateOfBirth = 1981

	//assigning function to a variable using short decleration 
	calcAge := func (DOB int) int{
		return 2021 - DOB
	}
	
	age := calcAge(dateOfBirth)
	
	fmt.Println("Firstname :",firstname, "Lastname :",lastname, "Date of Birth:", dateOfBirth, "Age", age)
}