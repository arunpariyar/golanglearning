package main

import (
	"encoding/json"
	"fmt"
)

type animal struct {
		Name string `json:"Name"`
	}

func main() {

	
	jsonBlob := `[{"Name":"Dog"},{"Name":"Cat"}]`

	jString := []byte(jsonBlob)

	var animals []animal

	err := json.Unmarshal(jString, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(animals)

}
