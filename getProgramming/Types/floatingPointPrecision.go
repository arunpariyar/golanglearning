package main

import "fmt"

func main(){
	var num float64
	fmt.Println(num)
	num = 1.0/3
	fmt.Println(num)
	//formating precision 
	fmt.Printf("%f\n", num)
	// 4width 2 precision
	fmt.Printf("%.4f \n", num)

}