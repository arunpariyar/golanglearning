//how long does it take to get to Mars
package main

import "fmt"

func main(){
	const lightSpeed = 299792

	var distance = 56000000

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000

	fmt.Println(distance/lightSpeed, "seconds")
	
}