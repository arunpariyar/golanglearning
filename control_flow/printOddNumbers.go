package main

import "fmt"

func main() {
	//use for loop to count till 100
	for i := 0; i <= 100; i++ {
		remainder := i % 2
		if remainder > 0 {
			fmt.Println(i)
		}
		continue
	}
	// use modulo to check if there is reminder is yes print otherwise continue
}
