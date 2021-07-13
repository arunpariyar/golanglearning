## **Short Variable Decleration**

The short variable decleration allows us to declare and assign a value to an identfier without us explicity defining the types. 

short decleration operator is := 

### GoLand Specification Definition

It is shorthand for a regular variable declaration with initializer expressions but no types:

`ShortVarDecl = IdentifierList ":=" ExpressionList`

Example 

`firstName := "James"`

`secondName := "Bond"`

we can create several variable at once as well using the short variable declerations

`firstName, secondName = "James", "Bond"`

## Some Basics

Some basic concepts that needs to be understood are identifiers, expression and statements. 

- Identifiers - Identifiers name program entities such as variables and types

- expression - An expression specifies the computation of a value by applying operators and functions to operands.

- statement - statement commands the computer to do something. statement comprise of expressions for the value a simple statement

firstName := "James" this tell the computer to create a variable with an identifier "firstName" and askes to make use of the short decleration variable to assign the value "James" to it.