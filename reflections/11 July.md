Today started with the Tod Mcleods, Golang course. 

I have gone through some portions of the couse before but i wanted to started from the very beginning so that's what i did today. 

**Important Keywords**

- package
- go modules 
- variadic parameter
- go workspace
- indentifier 

**package**

package refers to code that has been organised together that can be reused. In programming languages like javascript it is refered to as a Library. using package we are able to organise our code in a much better way. Also package can also refer to external code that we can integrate in our program for example go standard library are all packages that we can import and use in our program. for example the fmt package.

**Go modules**

Go Modules is the package management system or package manager used by GO. 

**Variadic Parameter **

in go we have functions that can take any type of input in unlimited amount. these inputs are simply refered to as variadic parameter and in the go documentation its refered to using ...interface{}. For example

func Println(a ...interface{}) (n int, err error)

Here the Prinln function can take in as many and any type of arguments. 

**Go workspace**
Prior to Go 1.11 go workspace was the main way of organising a go project. It comprised of three directories src, bin, pkg where src was used to keep all your  go code in packages, bin stored all the binary file and pkg stored all of the package archives. 

**Indentifer**

Indentifiers are names that we give to variable, type and functions