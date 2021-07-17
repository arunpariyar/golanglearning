// if we are using a struct just once and only once in our program then it is possible to substitue using anonymous struct... in all others case it better to use regular structs

package main

import "fmt"

func main(){
	// type dog struct{
	// 	name string 
	// 	breed string 
	// }

	// here put the composite literal directly into the variable and created the variable
	d1 := struct{ 
		name string 
		breed string 
	} {
		name : "Doug",
		breed: "pug",
	}

	fmt.Println(d1.name)
	fmt.Println(d1.breed)
}