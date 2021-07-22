// Write a program that finds the smallest number in this list:
// x := []int{ 48,96,86,68,
//         57,82,63,70,
//         37,34,83,27,
//         19,97, 9,17,
// }

package main

import "fmt"

func main() {
	x := []int{96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17, 1}
	//create a minimum value and assign it to the initial value of x
	smallestNo := x[0]

	for _, value := range x { 
		if smallestNo > value{
			smallestNo = value
		}
	}
	fmt.Println(smallestNo)

}

// the crucical thing about this program is that we must assign a value from inside the slice to be compared
