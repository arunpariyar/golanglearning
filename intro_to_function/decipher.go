package main

import "fmt"

func main() {
	// slice of int called values
	values := []int{3, 6, 9}
	// using the divisible by 3 function and and storing the values
	resultDivisibleByThree := divisibleBy3(sum, values...)
	fmt.Println(resultDivisibleByThree)
}
// the sum function that takes in varadic parameters of type int and returns an integer as well
func sum(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

func divisibleBy3(task func(sorted ...int) int, unsorted ...int) int {
	total := 0
	var sortedList []int

	for _, value := range unsorted {
		if value%3 == 0 {
			sortedList = append(sortedList, value)
		}
	}
	total = task(sortedList...)
	return total
}
