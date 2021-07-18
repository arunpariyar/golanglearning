package main

import "fmt"

func main(){
	operation, result1 := task("multiply", 1,2,3,4,5,6);
	fmt.Printf(" %v \t %v \n", operation, result1)
}

func task(do string, values ...int) (string, int) {
	total := 1
	switch true {
	case do == "multiply":
		for _, value := range values{
			total *= value
		}
	case do == "add":
		for _, value := range values{
			total += value
		}
	}

	return do, total
}