package main

import "fmt"

var prayer string

func main(){
	prayer = "So Hum"
	maximumValue := 5
	counter := 0
	for {	

		counter ++
		if counter > maximumValue {
			break
		}
		fmt.Println(prayer)

	}
}