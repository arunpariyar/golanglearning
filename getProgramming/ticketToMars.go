//program that generates random tickets to mars
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// create all the variables
var (
	spaceline, tripType string 
	days,price int 
)

func main(){
	
	

	fmt.Printf("%-17v %-4v %-10v %-5v \n", "Spaceline","Days","Trip Type","Price")
	fmt.Println("=======================================")

	for i := 0; i <= 10; i++ {
			//seed for the random number generator
	seed1 := rand.NewSource(time.Now().UnixNano())
	rand1 := rand.New(seed1)
	rand2 := rand.New(seed1)
	rand3 := rand.New(seed1)
	rand4 := rand.New(seed1)

	//randomNumber for spaceline
 	randomNumber := rand1.Intn(3) + 1

	//switch case for random number to decide airlines
	switch true {
	case randomNumber == 1:
		spaceline = "Space Adventures"
	case randomNumber == 2:
		spaceline = "Virgin Galactic"
	case randomNumber == 3:
		spaceline = "SpaceX"
	}
	//randomly choose speed from 16 to 30 km 
	speedMin := 16
	speedMax := 30

	//calculate randomSpeed
	randomSpeed := rand2.Intn((speedMax - speedMin + 1) + speedMax) + 1

	//generate number of days form the speed calculated
	//all constant values
	const distance = 62100000

	//calculate days
	seconds := distance/randomSpeed
	minutes := seconds/60
	hours := minutes/60 
	days := hours/24

	//randomly decide oneway for twoways deciding from if the first random number was odd or even
	randomReturn := rand3.Intn(2)+1
	if randomReturn == 1 {
		tripType = "One Way "
	} else {
		tripType = "Two Way"
	}
	//randomly generate price between 36 to 50 
	minPrice := 36
	maxPrice := 50
	randomPrice := rand4.Intn((maxPrice - minPrice + 1 ) + minPrice)
	// if even double the price
	if tripType == "One Way"{
		price = randomPrice
	} else {
		price = randomPrice * 2 
	}
	// fmt.Printf("%-17v %-4v %-10v %-5v \n", "Spaceline","Days","Trip Type","Price")
	fmt.Printf("%-17v %-4v %-10v %-5v \n",spaceline, days, tripType, price)
	}	
}