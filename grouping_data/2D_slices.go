package main

import "fmt"

func main(){
	//array and slice are one dimensional but if is possible to create two dimensional array and slices

	var OneDArray [2]int
	var TwoDArray [2][2]int
	
	OneDArray = [2]int{1,2}
	TwoDArray = [2][2]int {{1,2},{1,2}}

	OneDSlice := []int{1,2}
	TwoDSlice := [][]int{{1,2},{3,4}}

	fmt.Println("**** One and Two Dimension Array ****")
	fmt.Println(OneDArray)
	fmt.Println(TwoDArray)

	fmt.Println("**** One and Two Dimension Slice ****")
	fmt.Println(OneDSlice)
	fmt.Println(TwoDSlice)


}