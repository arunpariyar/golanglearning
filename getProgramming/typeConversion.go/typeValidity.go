package main

import (
	"fmt"
	"math"
)

func main(){
	var value16 uint16
	var value uint16 = 15 

	fmt.Printf("%T \n", value)

	if value > 0 && value < math.MaxUint16 {
		value16 = uint16(value)
	}

	fmt.Println(math.MaxUint16)
	fmt.Printf("%v %[1]T", value16)
}