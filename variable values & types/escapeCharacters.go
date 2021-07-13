package main

import "fmt"

func main(){

	var para = fmt.Sprint("This is a quick test to print all the escape character. The first is the alert or the bell \a . Now the backspace \b . Now the form feed \f . newline of line feed \n and of course the carrige return or enter \r. Moving on the horizontal tab \t and then the vertical tab \v. Finally we have some escape characters backlash \\. the single quote and the double quote\".And that is all of the escape charaters ")

	fmt.Println(para)
}