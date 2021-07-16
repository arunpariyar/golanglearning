package main

import "fmt"

func main(){
	type tree struct{
		name string
		family string
		age int
	}
	
	type park struct{
		name string
		tree
		noOfTrees int
		area int
	}

tierGarten := park {
	name: "Tier Garten",
	tree: tree {"oak", "oakFamily", 350},
	noOfTrees: 250000,
	area: 210,
}

fmt.Printf("%v coveres the area of %v hectre and has around %v trees. That are mainly %v of %v family the oldest tree is %v years old", tierGarten.name, tierGarten.area, tierGarten.noOfTrees, tierGarten.tree.name, tierGarten.family, tierGarten.age )
}
