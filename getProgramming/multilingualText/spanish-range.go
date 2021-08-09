package main

import (
	"fmt"
)

func main(){
	question := "¿Cómo estás?"
	
	for index, value := range question{
		fmt.Printf("%v  %c \n", index, value )
	}
}