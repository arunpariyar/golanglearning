// (C = (F − 32) * 5/9).

package main

import "fmt"

func main(){
	fahrenherit := 120 

	celcius := toCelcius(120)

	fmt.Printf("%v Fahernheit in celcius is %v \n", fahrenherit, celcius)
}

func toCelcius(value int) int {
	converted := (value - 32) * 5/9

	return converted
}