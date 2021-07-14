package main

import "fmt"

var prayer string
var blessing string

func main() {
	prayer = "Om Mane Padme Hum ༅"
	blessing = "May all being be happy, May all beings be peaceful"
	for i := 0; i <= 5; i++ {
		fmt.Printf("\t %v %s \n", i, prayer)
		for j := 0; j <= 10; j++ {
			fmt.Printf("%v %s \n", j, blessing)
		}
	}
}
