// Write a function with one variadic parameter that finds the greatest number in a list of numbers.

package main

import "fmt"


func main(){
	result := calcGreatestNumber(6,7,8,9,23,26,87)
	fmt.Println(result)
}

func calcGreatestNumber(values ...int) int{
	var greatestNumber int
	greatestNumber = values[0]
	for _, value := range values{
		if greatestNumber < value {
			greatestNumber = value
		}
	}
	return greatestNumber
}