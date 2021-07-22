// Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers

package main

import "fmt"

func main(){
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}

//created the fucntion that returns a func a unsigned int
func makeOddGenerator() func() uint{

	//setting the value of i to be 1
	i := uint(1)
	// creating an anyonomus return function that takes no argument but returns an uint ret
	return func () uint  {
		ret := i
		i += 2
		return ret 
	}

}

// this program is also a great example of closure where variable i keeps its state everytime the nextodd() function is called successively