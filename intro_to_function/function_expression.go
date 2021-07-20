//when we assign a function to a variable its is called a function expression

package main

import "fmt"


func main(){
	// we are saving the function in to funcExpression and running it later with the variable name followed by funcExpression()
	funcExpression := func(){
		fmt.Println("I am a function expression")
	}
	funcExpression()

	//creating a sum function using function expression 
}

