package main

import "fmt"

func main() {
	at := '@'
	var smiley = '😁'
	var acute rune = 'é'

	fmt.Printf("%c = %[1]v, %c = %[2]v, %c = %[3]v \n",at, smiley, acute)
}