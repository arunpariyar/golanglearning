package main

import "fmt"

func main() {
	result := fibonacci(10)
	fmt.Println(result)
}
func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}