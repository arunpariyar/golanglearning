// variadic parameter written as ... refers to unlimited numbers of values of the specified type

package main

import "fmt"

func main() {
	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sum)
}

// creating a function that takes in a varadic parameter

func sum(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
