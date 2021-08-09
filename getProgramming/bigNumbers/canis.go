// Canis Major Dwarf is the closest known galaxy to Earth at 236,000,000,000,000,000 km from our Sun (though some dispute that it is a galaxy). Use constants to convert this dis- tance to light years.

package main

import "fmt"

func main(){
	const distance = 236000000000000000
	const lightSpeed = 299792

	lightYears := distance / lightSpeed / (60*60*24) / 365

	fmt.Println("Canis Major Dwarf is ", lightYears ," Light Years away from Earth")
}