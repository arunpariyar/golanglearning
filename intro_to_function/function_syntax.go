package main

import "fmt"

func main() {
	// the standard syntax of a function is func (r receiver )(parameter)(return){code block}

	//calling the sayHello function
	sayHello()
	sayHelloWithName("John")

	//we save the return for the function to a variable line and then print it later
	line := sayHelloWithReturn("John", "Doe")
	fmt.Println(line)

	//calling the function and saving the returns to variables
	greeting, age := sayHelloWithMulitpleReturns("John","Doe",1990)
	fmt.Println(greeting)
	fmt.Printf("I am %v years old \n", age)
}

//creating a function without return that says hello world
func sayHello() {
	fmt.Println("Hello World")
}

//creating a function that takes a name as a parameter but doesnt return anything apart form printing a hello message
func sayHelloWithName(name string){
	fmt.Printf("Hello  %v \n",name )
}

//creating a function that takes two parameter and returns a single return of type string 
func sayHelloWithReturn (firstname, lastname string)(string){
		line := fmt.Sprintf("Hello %s %s", firstname,  lastname)
	 return line
}

//create a function that take parameters and two returns 
func sayHelloWithMulitpleReturns(firstname, lastname string ,yearOfBirth int )(string, int){
	greeting := fmt.Sprintf("Hey i am %s %s ", firstname, lastname)
	age := 2021 - yearOfBirth
	// creating multple returns
	return greeting, age 
}
