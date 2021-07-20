package main

import "fmt"

type person struct{ 
	name string 
	dateOfBirth int
}

func (person person) age() int{
	age := 2021 - person.dateOfBirth
	return age
}

func main(){
	p1 := person {
		name: "john",
		dateOfBirth: 1990,
	}
	age:=p1.age();

	fmt.Println(age)
	//Anoynomus (self-executing)function that prints 
	func (p person, age int) {
		fmt.Printf( "Hey i am %v and i am %v years old \n", p.name, age)
	}(p1, age)

}