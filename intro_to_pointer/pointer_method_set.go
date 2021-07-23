package main

import "fmt"

type person struct { 
	firstname string
	lastname string
}

func (p person) says(){
	fmt.Println("Hey I am human")
}

func main(){
	p1 := person{
		firstname: "Harry",
		lastname: "Porter",
	}
	p2 := &p1
	introduce(*p2)

}
