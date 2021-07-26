package main

import (
	"encoding/json"
	"fmt"
)

// the name fields must be capital for it to work
type animal struct {
	Name   string
	Family string
}

func main() {
	a1 := animal{
		Name:   "dog",
		Family: "wolf",
	}

	a2 := animal{
		Name:   "sparrow",
		Family: "bird",
	}

	allAnimals := []animal{a1, a2}

	allAnimalsJson, err := json.Marshal(allAnimals)

	if err != nil {
		fmt.Println("error :", err)
	}

	fmt.Println(string(allAnimalsJson))

}
