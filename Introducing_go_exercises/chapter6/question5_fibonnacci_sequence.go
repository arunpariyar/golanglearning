// The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).

package main

import "fmt"

func main() {

	result := fibonacci(10)
	fmt.Println(result)

}
func fibonacci(value int) int {
	if value == 0 {
		return 1
	} else if value == 1{
		return 1
	} else {
		return fibonacci(value - 1) + fibonacci( value - 2 )
	}	
}

// This program is not correct will have to look at it again later 
