package main

import "fmt"

type person struct{
	firstname string
	lastname string 
	dateOfBirth int
}

type dog struct { 
	name string
	dateOfBirth int
}

type animal interface { 
	calcAge() int
}

func main(){

	p1 := person { 
		firstname: "John",
		lastname: "Smith",
		dateOfBirth: 1980,
	}

	d1 := dog { 
		name: "Milo",
		dateOfBirth: 2005,
	}

	age(p1)
	age(d1)
	fmt.Println(greet(p1))
	fmt.Println(greet(d1))
}

func age(a animal)  { 
	fmt.Println(a.calcAge())
}

func greet(a animal) string{
	age := a.calcAge()
	return fmt.Println("my age", age)
}
func (d dog) calcAge() int{
	return 2021 - d.dateOfBirth
}

func (p person) calcAge() int {
	return 2021 - p.dateOfBirth
}