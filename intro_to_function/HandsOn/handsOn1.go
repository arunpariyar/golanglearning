package main

import "fmt"


func main(){
	// call both funcs
	minChanting := foo()
	number, sentence := bar()

	// print out their results
	fmt.Println("The minimum chanting rounds is", minChanting)
	fmt.Println(sentence, number)
}
// create a func with the identifier foo that returns an int
func foo() int{
	return 108
}
// create a func with the identifier bar that returns an int and a string
func bar() (int, string){
	return 108, "Minimum Chanting Number"
}