package main

import "fmt"

func main(){
	DOB := 1990
	for {
		if DOB == 2021 { 
			break
		}
		
		fmt.Printf("%v \n", DOB)
		DOB++
	}
}