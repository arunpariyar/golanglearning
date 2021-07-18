//if we want a function to execute a the very end of a code block or program irregradles of where the function is called we can used defer.

package main

import "fmt"



func main(){
	
	defer closefile()
	doSomething()
	anotherTask()

}

func doSomething(){
	fmt.Println("I am doing something with a file")
}

func anotherTask(){
	fmt.Println("I am also working with a file")
}

func closefile(){
	fmt.Println("I am now closing the file to avoid the program using it any further")
}