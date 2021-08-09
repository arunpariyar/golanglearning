package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "shalom"

	length := []rune(word)

	fmt.Println("word has the length of", len(length))
	
	for i:=0; i < utf8.RuneCountInString(word); i++{
		c:= word[i]
		fmt.Printf("%c\n",c)
	}

	c := 'g'
	c = c - 'a'+'A'
	fmt.Printf("%c \n",c)
}
