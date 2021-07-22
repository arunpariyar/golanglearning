// sum is a function that takes a slice of numbers and adds them together. What would its function signature look like in Go?

package main

import "fmt"

func main(){
	x:= []int{1,2,3,4}
	result := sum(x)
	fmt.Println(result)

}
func sum(values []int) int{
	total := 0
	for _, value := range values{
		total += value
	}
	return total
}