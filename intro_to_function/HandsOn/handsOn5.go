/* create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type square struct {
	length float32
}

type shape interface {
	area() float32
}

func main() {
	c1 := circle{
		radius: 2.5,
	}
	s1 := square{
		length: 5,
	}

	info(c1)
	info(s1)

}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float32 {
	return s.length * s.length
}

func info(sh shape) {
	switch sh.(type) {
	case circle:
		fmt.Println("The area of circle is", sh.(circle).area())
	case square:
		fmt.Println("The area of circle is", sh.(square).area())
	}
}
