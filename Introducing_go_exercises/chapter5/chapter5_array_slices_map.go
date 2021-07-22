// How do you access the fourth element of an array or slice?
//using the name of the array[3]

//What is the length of a slice created using make([]int, 3, 9)?
//length of the slice is 3 ... here []int -> slice of integers, 3 -> length and 9-> capacity of the slice
// Given the following array, what would x[2:5] give you?
// x := [6]string{"a","b","c","d","e","f"}
//[c,d,e]

package main

import "fmt"

func main(){
	numbers := make([]int,3,9)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
}