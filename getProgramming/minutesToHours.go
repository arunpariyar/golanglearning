package main

import "fmt"

type minutes float64
type hours float64

func main(){
	var m minutes = 2343
	var h hours= minutesToHours(m)
	
	fmt.Printf("%v minutes is %v hours \n", m, h)

}

func minutesToHours(m minutes) hours { 
	return hours(m / 60)
}