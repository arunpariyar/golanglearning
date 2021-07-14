package main

import "fmt"

func main() {
	maximumValue := 100
	counter := 0

	for counter < maximumValue {
		counter++
		fmt.Println("The value of counter is now:", counter)
	}

}