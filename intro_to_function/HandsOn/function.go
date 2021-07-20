package main

import "fmt"

func main(){
	message := inceptionLevel1()()
	fmt.Println(message)
}

func inceptionLevel1() func()string{
	return func() string{
		return fmt.Sprintln("what is reality")
	}
}