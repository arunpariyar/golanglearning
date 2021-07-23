package main

import "fmt"

func main(){
	//allocating a value to a variable in our computer means that we are storing it in the computers memory somewhere.
	houseNo := 33 
	
	//if I want to know the address of that value I can use & to get it
	fmt.Println(&houseNo)
	
	//but also if I want to know the value a memory is pointing to I can use
	fmt.Println(*&houseNo)
	
	//but what is the underlying type of the pointer
	fmt.Printf("%T \n", &houseNo) // this prints *int the pointer
	
	//it is also possible to save the address of the variable 
	addressOfHouseNo := &houseNo
	
	//Curious what is the type of addressOfHouse
	fmt.Printf("%T \n",addressOfHouseNo)
	
	//and if i wanted to print the value stored at the address variable
	fmt.Println(*addressOfHouseNo)
	
	//the above is also the same as 
	fmt.Println(*&houseNo)

	//if I use the *&variablename - I am derefrencing the value 
}