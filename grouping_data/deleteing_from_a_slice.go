package main

import "fmt"


func main(){
	friends := []string{"John","Jake","Jina","Jack"}
	// keep only jake and jina as friends
	friends = append(friends[:0], friends[2:]...)
	fmt.Println(friends)
}