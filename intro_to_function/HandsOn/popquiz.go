package main

import "fmt"

type Toy struct {
	Color string
}

func main() {
	toys := []Toy{
		{Color: "red"},
		{Color: "green"},
	}
	fmt.Println(toys)
	for _, toy := range toys {
		if toy.Color == "green" {
			toy.Color = "blue" // why does this value not persist
		}
	}
	fmt.Println(toys)
}
