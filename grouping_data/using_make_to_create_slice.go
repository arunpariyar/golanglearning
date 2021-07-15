package main

import "fmt"

func main(){
	//slice are built on top of an array if we know what will be be capacity of our size and the current length we can create a slice using make that defines the length and capacity of our slice. 
	
	//i want to create a slice of length 2 that has the capacity to grow till 10. 
	// the syntax is make(type, length, capacity)

	friends := make([]string, 3, 10)

	friends[0] = "Jamie"
	friends[1] = "Jil"
	

	fmt.Println(friends)
	fmt.Println(len(friends))
	fmt.Println(cap(friends))
	// it is however possible to append the friends and add new elements that excede the current length
	fmt.Println("**** adding jack *****")
	friends = append(friends, "Jack", "Jinny", "Jerry", "Jackie")

	fmt.Println(friends)
	fmt.Println(len(friends))
	fmt.Println(cap(friends))
}