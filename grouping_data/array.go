package main

import "fmt"

func main(){
	var friends [5]string 
	friends[0] = "arun"
	friends[1] = "me"
	friends[2] = "myself"
	friends[3] = "god"

	for i:=0 ; i<len(friends); i++{
	 fmt.Println(friends[i])
	}
	
}