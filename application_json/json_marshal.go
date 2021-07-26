package main

import (
	"encoding/json"
	"fmt"
)

type superhero struct {
	Name  string
	Alias string
	Brand string
}

func main() {
	s1 := superhero{
		Name:  "spider-man",
		Alias: "peter parker",
		Brand: "marvel",
	}
	s2 := superhero{
		Name:  "super-man",
		Alias: "Clark Kent",
		Brand: "DC",
	}

	allHeroes := []superhero{
		s1, s2,
	}
	fmt.Println("Struct:",allHeroes)
	
	allHeroesJson, err := json.Marshal(allHeroes)
	
	if err != nil { 
		fmt.Println("err", err)
	}
	fmt.Println("JSON:",string(allHeroesJson))
}