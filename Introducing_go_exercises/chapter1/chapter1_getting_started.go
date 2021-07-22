// 1. What is whitespace?
/* 	Newlines, spaces and tabs are white space becuase we can't see them. the langague doesnt care about the white space it used for readablitiy reasons */

//2. What is a comment? What are the two ways of writing a comment?
/*  comments are used to store information about our code for documentaton purpose. single line comment start with "//" while multiline comments are written inside /* */
// 3. Our program began with package main. What would the files in the fmt package begin with?
// They would begin with package fmt
//4. We used the Println function defined in the fmt package. If you wanted to use the Exit function from the os package, what would you need to do?
// we would import the os package and invoke the Exit function using os.Exit()

//5. Modify the program we wrote so that instead of printing Hello, World it prints Hello, my name is followed by your name.

package main

func main(){
	name:= "John Doe"

	println("Hello, ", name )
}