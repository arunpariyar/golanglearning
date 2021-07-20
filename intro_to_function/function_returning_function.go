package main

import "fmt"

func main() {
	
	message := level1()()

	fmt.Println(message)

}

func level1() func() string {
	return func() string{
		return fmt.Sprintln("Hello from level 0")
	}
}
