// pass a func into a func as an argument
package main

import "fmt"

func main() {
	message := sayHello(hello,` Jina`)
	fmt.Println(message)

 }

func sayHello(task func() string, name string) string {
		message := task()
		return message + name
	}

func hello() string {
	return "Hello"
}

// Here I pass a function as an argument and call it inside the code block.


//Where exactly is the callback ?