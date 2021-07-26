package main

import (
	"encoding/json"
	"fmt"
)

type superhero struct {
	Name string `json: "Name"`
}

func main() {
	jsonString := `[{"Name":"Spider-man"},{"Name":"Superman"}]`
	byteString := []byte(jsonString)

	var heroes []superhero

	err := json.Unmarshal(byteString, &heroes)
	if err != nil { 
		fmt.Println("err", err)
	}

	fmt.Println(heroes)

}