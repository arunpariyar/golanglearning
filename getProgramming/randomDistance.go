// generate a random distance from 56000000 to 401000000km

// generate random num
package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	min := 56000000
	max := 401000000

	fmt.Println(r1.Intn(max - min +1) + min )
}
