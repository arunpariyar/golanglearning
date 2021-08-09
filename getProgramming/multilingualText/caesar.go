package main

import "fmt"

func main() {
	quote := "L fdph L vdz L frqtxhuhg."

	for i := 0; i < len([]rune(quote)); i++{
		char := quote[i] - 3
		
		fmt.Printf("%c ", char)
	}
}