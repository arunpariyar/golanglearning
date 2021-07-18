// Hands-on exercise #1
// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

package main

import "fmt"

type person struct {
	firstname string
	secondname string
	favflavours []string
}

func main(){
	p1 := person{
		firstname: "Jamie",
		secondname: "Goldburn",
		favflavours: []string{"vanilla","cherry"},
	}
	p2 := person{
		firstname: "Jack",
		secondname: "Spades",
		favflavours: []string{"metal","gunpower"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
