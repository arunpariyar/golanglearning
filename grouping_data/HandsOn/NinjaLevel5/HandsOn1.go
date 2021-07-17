package main

import "fmt"

func main(){
	type person struct{
		first_name string
		last_name string 
		fav_icream_flavours []string
	}

	p1 := person {
		first_name: "Harry",
		last_name: "Potter",
		fav_icream_flavours: []string{"Chocolate","Strawberry"},
	}

	fmt.Println(p1)
}