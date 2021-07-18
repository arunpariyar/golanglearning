package main

import "fmt"

func main() {
	// basic syntax
	type dog struct {
		name  string
		breed string
		age   int
	}

	type car struct{ 
		name string 
		brand string 
		model string
		year int 
	}

	// create a structure of type dog
	d1 := dog {
		name:  "Annie",
		breed: "Japanes Chino",
		age:   9,
	}

	fmt.Println(d1)
	fmt.Println(d1.name)
	fmt.Println(d1.breed)
	fmt.Println(d1.age)

	//creating a structure of type car
	whiteStar := car { 
		name: "White Star",
		brand: "Tesla",
		model: "S",
		year: 2012,
	} 

	fmt.Printf("The %v model %v car from %v became the first electric car to top the monthly new-car-sales ranking in any country.", whiteStar.name, whiteStar.model, whiteStar.brand )
	
}
