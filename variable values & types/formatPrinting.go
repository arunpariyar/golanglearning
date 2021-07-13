package main

import "fmt"

func main(){
	var name = "James"

	var p = fmt.Sprintf("%s\t %q\t %x %X\t",name,name,name,name )

	fmt.Println(p)
}