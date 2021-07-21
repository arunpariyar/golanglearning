// 2. Write a program that prints out all the numbers between 1 and 100 that are evenly divisible by 3 (i.e., 3, 6, 9, etc.).

package main

import "fmt"

func main(){

	//use for loop 
	//inside if check, if yes print or else continue
	for i:= 0; i<=100; i++{
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
