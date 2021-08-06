//my weight loss program.
package main

import (
	"fmt"
	"math"
)

//the main function
func main(){
	var myWeight float32 = 63

	weightDifference := myWeight * 0.38

	weightInMars := myWeight - weightDifference

	weightRounded := math.RoundToEven(float64(weightInMars))

	fmt.Printf("My weight in mars is %v \n", weightRounded )

}