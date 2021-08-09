// Cipher the Spanish message “Hola Estación Espacial Internacional” with ROT13. Mod- ify listing 9.7 to use the range keyword. Now when you use ROT13 on Spanish text, char- acters with accents are preserved.

package main

import "fmt"

func main() {
	message := "Hola Estación Espacial Internacional"

	for _, value := range message {
		if value > 'a' && value < 'z' {
			value = value + 13
			if value > 'z' {
				value -= 26
			}
		} else if value > 'A' && value < 'Z' {
			value += 13
			if value > 'Z' {
				value -= 26
			}
		}
		fmt.Printf("%c", value)
	}
}
