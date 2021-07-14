package main

import "fmt"

func main(){
	DOB := 1990
	ThisYear := 2021
	for DOB <= ThisYear {
		fmt.Printf("%v \n", DOB)
		DOB++
	}
}