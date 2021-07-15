package main

import (
	"fmt"
	"reflect"
)

func main(){
	var num []int 

	num = []int{1, 2,3,4,5}

	fmt.Println(num)
	fmt.Println(reflect.TypeOf(num))
}