package main

import "fmt"

func main() {
	type player struct {
		name       string
		experience int
	}
	type coach struct {
		name     string
		nickname string
	}

	//team has embedded struct within it
	type team struct {
		player
		coach 
	}

	griffindor := team{
		player: player {"Harry",3},
		coach: coach {"sergay","sir"},
	}

	fmt.Println(griffindor)
}
