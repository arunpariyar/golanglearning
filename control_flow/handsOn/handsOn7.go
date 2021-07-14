package main

import "fmt"

func main(){
	haveCard := true
	tempAccess := true
	if haveCard || tempAccess{ 
		fmt.Println("You may enter the building")
	} else if tempAccess{ 
		fmt.Println("You may enter the building")
	} else { 
		fmt.Println("We are very sorry but with the card or the temporary access you are not allowed to enter the building")
	}
}