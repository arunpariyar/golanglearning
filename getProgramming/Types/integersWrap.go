package main

import "fmt"

func main() {
	var red uint8 = 0
	red--
	fmt.Println(red) // prints 255

	var number int8 = -128
	number--
	fmt.Println(number)

	var num uint16 = 65535
	num++
	fmt.Println(num)
}