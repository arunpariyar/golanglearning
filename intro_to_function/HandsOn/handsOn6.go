// Build and use an anonymous func

package main

import "fmt"

func main(){

	sum := func (x int, y int ) int{
		return x + y
	}

	fmt.Println(sum(1,2))
}

func add(x int, y int) int{
	return x+y
}