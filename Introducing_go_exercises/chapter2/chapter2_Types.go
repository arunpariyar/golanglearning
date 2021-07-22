// 1. How are integers stored on a computer?
//integers are stored on computer using base-2 binary system
/* 2. We know that (in base 10) the largest one-digit number is 9 and the largest two-digit number is 99. Given that in binary the largest two-digit number is 11 (3), the largest three-digit number is 111 (7) and the largest four-digit number is 1111 (15), what’s the largest eight-digit number? (Hint: 101−1 = 9 and 102−1 = 99) */
//11111111 => 255

// 4. What is a string? How do you find its length?
//string is a sequence of bytes we can find the lenth using len(string) function
// 5. What’s the value of the expression (true && false) || (false && true) || !(false && false )
// true

// 3. Although overpowered for the task, you can use Go as a calculator. Write a pro‐ gram that computes 32,132 × 42,452 and prints it to the terminal (use the * oper‐ ator for multiplication).

package main

import "fmt"

func main() {

	fmt.Println(len("Hello"))
	result := multiply(32132, 42452)

	fmt.Println(result)
}

func multiply(values ...int) int{
	total := 1
	for _, value := range values {
		total *= value
	}
	return total
}
