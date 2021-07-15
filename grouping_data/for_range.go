package main

import "fmt"

func main(){
	//composite literal to create a slice
	captainPlanet := []string{"earth","water","fire","air"}

	for index, value := range captainPlanet{
		fmt.Println(index, value)
		
	}
	fmt.Println("By Your powers combined. I am captain planet")
}