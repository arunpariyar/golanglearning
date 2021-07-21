// (1 ft = 0.3048 m)

package main

import "fmt"

func main() {
	fmt.Println("Please enter distance in feet: ")
	var input float32
	fmt.Scanf("%f", &input)
	inMeter := toMeter(input)

	fmt.Printf("%v feet in meter is %v \n", input, inMeter)
}

func toMeter(inFeet float32) float32 {
	return inFeet * 0.3048
}
