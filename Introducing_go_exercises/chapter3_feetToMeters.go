// (1 ft = 0.3048 m)

package main

import "fmt"

func main() {
	feet := 100

	inMeter := toMeter(float32(feet))

	fmt.Printf("%v feet in meter is %v \n", feet, inMeter)
}

func toMeter(inFeet float32) float32 {
	return inFeet * 0.3048
}
