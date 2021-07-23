// Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable

package main

import "fmt"

func main(){
	number := nextOddNumber()
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
}

func nextOddNumber() func() int{
	start := 1
	return func() int {
			start += 2
			return start
	}

}