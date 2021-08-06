// generating random numbers in go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	num := genRand(56)
	fmt.Println(num)
}

func genRand(numrange int) int{
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(numrange) + 1
}