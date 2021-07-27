package main

import (
	"encoding/json"
	"fmt"
)

// create the struct type for the data that will be converted to json
type Agent struct{
	First string `json:"First"`
	Last string `json:"Last"`
	Age int `json:"Age"`
	Sayings []string `json:"Sayings"`
}
func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	
	// convert the string literal to []byte
	stringByte := []byte(s)

	// create a variable of the defined struct type
	var allAgents []Agent
	
	err := json.Unmarshal(stringByte, &allAgents)

	if err != nil {
		fmt.Println("error:",err)
	}

	for index, agent := range allAgents{
		fmt.Println("Agent No", index)
		fmt.Println(agent.First)
		fmt.Println(agent.Last)
		for _, saying := range agent.Sayings{
			fmt.Printf("\t %v\n", saying)
		}
	}
}
