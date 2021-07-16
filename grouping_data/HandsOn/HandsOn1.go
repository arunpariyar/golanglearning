package main

import "fmt"

func main() {
	arrayOfInt := [5]int{1, 2, 3, 4, 5}

	for _, value := range arrayOfInt {
		fmt.Printf("%v is of type %T\n", value, value)
	}
}
