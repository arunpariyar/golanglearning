package main

import "fmt"

//type decleration
type hours float64
type minutes float64
type seconds float64

func main(){

	var h1 hours = 2
	var m1 minutes = 2 
	
	fmt.Println( h1," Hours in seconds", h1.toSeconds())
	fmt.Println( m1, " Minutes in seconds",m1.toSeconds())

}

//method decleration keyword receiver methodName result
func (h hours) toSeconds() seconds{
 return seconds(h * 60 * 60)
}

func (m minutes) toSeconds() seconds{
	return seconds(m * 60)
}