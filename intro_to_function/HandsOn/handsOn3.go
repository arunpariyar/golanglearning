//Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

package main

import "fmt"

func main(){
	defer sayNamaste() // this function will execute at the very end before the end of the block
	sayHello()
	sayGutenTag()
	fmt.Println("End of the code block")
}

func sayNamaste(){
	fmt.Println("Namaste")
}

func sayHello(){
	fmt.Println("Hello")
}

func sayGutenTag(){
	fmt.Println("hallo ! Guten tag")
}