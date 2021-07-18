// in go methods are created by attaching them to the type of the struct

package main

import "fmt"

type person struct {
	firstname   string
	lastname    string
	dateOfBirth int
}

func (p person) age() {
	fmt.Printf("Hey I am %v, I am %v years old.\n", p.firstname, 2021-p.dateOfBirth)
}

func (p person) sayHello() {
	fmt.Printf("Hey I am %v %v and I was born in %v. \n", p.firstname, p.lastname, p.dateOfBirth)
}

func main() {

	person1 := person{
		firstname:   "John",
		lastname:    "Doe",
		dateOfBirth: 1990,
	}

	person1.sayHello()
	person1.age()
}
