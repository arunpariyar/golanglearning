// When to use pointer is a good question ?
// Ideally, If we have work with big volume of data to reduce the workload on the computer pointer can be used. \
// a small demo however is given below.

package main

import "fmt"

func main(){
	age := 40
	fmt.Println(&age) //printing the address of houseNo
	passByPointer(&age) // passing to the function the pointer of age 
	fmt.Println(&age)

}

//the type of a memory address is a pointer * followed by data type is this case an integer
func passByPointer(value *int){
	*value = 15	// here i am changing the value of the pointer value
}