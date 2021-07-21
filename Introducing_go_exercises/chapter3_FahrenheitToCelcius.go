// (C = (F − 32) * 5/9).

package main

import "fmt"

func main(){
	fmt.Print("Please enter Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	celcius := toCelcius(120)
	fmt.Printf("%v Fahernheit in celcius is %v \n", input, celcius)
}

func toCelcius(value int) int {
	converted := (value - 32) * 5/9
	return converted
}