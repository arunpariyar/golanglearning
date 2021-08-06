package main

import "fmt"

//type age
type age int

func main() {
	var a1 age = 30
	fmt.Println(a1.birthYear())
}

//method birthYear for type age
func (a age) birthYear() int {
	return 2021 - int(a)
}
