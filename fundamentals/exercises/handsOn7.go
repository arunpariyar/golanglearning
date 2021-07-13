package main

import "fmt"

func main(){
	start := 2
	next := start << 1
	
	fmt.Printf("%d \t  %b\t %X \t \n", start, start, start)
	fmt.Printf("%d \t  %b\t %X \t \n", next, next, next)

}