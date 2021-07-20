package main

import "fmt"

//dog struct
type dog struct {
	name  string
	breed string
}

// cat struct
type cat struct {
	name  string
	breed string
}

//func receiver parameters return {codeblock}
func (dog dog) says() {
	fmt.Println(dog.breed, dog.name, "says woof woof")
}

func (cat cat) says() {
	fmt.Println(cat.breed, cat.name, "says meow meow")
}

//ANIMAL INTERFACE FOR THOSE WITH says() METHOD
type animal interface {
	says()
}

//ATTACHING THE METHOD INTRODUCE TO ANIMAL NOTE: when creating methods for interface the name comes first and then the interfaces name ... here the name introduce for (animal animal)
func introduce(animal animal) {
	switch animal.(type) {
	case dog:
		fmt.Println("Hey I am an animal", animal.(dog).name)
	case cat:
		fmt.Println("Hey I am an animal", animal.(cat).name)
	}

	animal.says()

}

func main() {

	dg := dog{
		name:  "Jackie",
		breed: "Dalmation",
	}

	ct := cat{
		name:  "Jenny",
		breed: "Gentle",
	}

	introduce(dg)

	introduce(ct)

}
