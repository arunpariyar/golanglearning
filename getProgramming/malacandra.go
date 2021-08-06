// determine how fast it needs to travel llto reach malacandra in 28 days

package main

import "fmt"

func main(){
	distance := 56000000
	days := 28 
	hours := 24

	fmt.Println(distance / (days * hours))
}