// how many days would it take to reach Mars

package main

import "fmt"

func main(){
	const speed = 100800
	var distance = 96300000

	time := distance/speed
	days := time / 24

	fmt.Println("It will take ", days," days to reach to Mars.")
}