package main

import (
	"fmt"
	"reflect"
)

func main(){
	
	fmt.Print("Please enter something:")
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println("You entered ", input)
	fmt.Println(reflect.TypeOf(input))
}