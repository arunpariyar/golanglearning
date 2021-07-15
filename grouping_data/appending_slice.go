package main

import "fmt"

func main() {
	friends1 := []string{"John", "Jake", "Jude"}
	friends2 := []string{"Lina", "Lisa", "Lydia"}

	allFriends := append(friends1, friends2...)

	fmt.Println(allFriends)
}