package main

import (
	"fmt"
	"math/big"
)

func main() {
	distance := new(big.Int)
	distance.SetString("24000000000000000000",10)

	speed := big.NewInt(299792)
	secondsPerDay := big.NewInt(60 * 60 *24)

	fmt.Println("The distance from Earth to Mars is", distance, "second per day ", secondsPerDay)

	seconds := new(big.Int)
	seconds.Div(distance, speed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay )

	fmt.Println("It will take", days, "days to reach mars at light speed")
}