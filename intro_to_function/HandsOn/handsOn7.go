// Assign a func to a variable, then call that func

package main

import "fmt"

func main() {

	multiply := func(values ...int) int {
		result := 1
		for _, value := range values {
			result *= value
		}
		return result
	}

	result := multiply([]int{1,2,3,4,5}...)
	fmt.Println(result)
}
