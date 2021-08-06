package main

import (
	"fmt"
	"strconv"
)

func main(){

	 text := "33.33"

	num, _ := strconv.ParseFloat(text, 64)

	fmt.Printf("%T \n", num)

}