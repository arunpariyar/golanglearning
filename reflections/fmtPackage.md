#Fmt package

fmt package is used for doing formatted I/O functions. which simply means it is used to do formatted printing.

for general use
%v - value in default format
%T - print the type of the variable

the Default format for various types are:

- bool: %t
- int, int8 etc: %d
- float32, complex64 etc: %d
- string: %s

###fmt functions

fmt.Print, fmt.Println, fmt.Printf : all related to printing to display.

fmt.Fprint, fmt.Fprintln, fmt.Fprintf: all related to printing to a file or webserver

fmt.Sprint, fmtSprintf, fmt.Sprintln: all relate to printing to a string value.
