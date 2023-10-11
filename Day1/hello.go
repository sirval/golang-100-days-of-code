package main

import "fmt"

// import

//TODO: VARIABLES
/*Variables in GoLang has 3 visibility Package level variables, scope (function) level and finally Global variables
<+-----------------------------------------------------------+>
---------------------------------------------------------------
scope or local variables will always take precedence over package visibility variables
E.G

var Account string = 'hello' //globally visible ---- see upper cased  first letter variable
var a string = 'hello' //package visible ---- when declared outside the scope within a package
func main(){
	var a string = 'hello' //scope visible ---- when declared inside a function
}

/*|----------------------------------------------------|
|----------------------------------------------------|*/

/*Example 1
func main() {
	fmt.Println("Welcome to Go programming")
	//variable declaration
	// var a int
	// a = 10

	//2 var a int = 12
	a := 15
	fmt.Println("c = ", a)
}*/

/*Example 2*/

/*var username string = "IamValentine"
var (
	name    string = "Ikenna Valentine"
	id      int32  = 12
	surname string = "Ohuka"
	age     int    = 28
)

func main() {
	fmt.Println("%v", "My name is "+surname+" "+name+" I am "+fmt.Sprint(age)+" years old")
}*/

//TODO: PRIMITIVES
/* |------------------------------------------------------------------|
There are majorly 3 primitive data types in go
Boolean
numeric
String
*/
func main() {
	isVerified := false
	num1 := 12
	num2 := 10
	fmt.Printf("%v, %T\n", isVerified, isVerified)
	//Bitwise operators
	fmt.Printf("%v\n", num1&num2)
	fmt.Printf("%v\n", num1|num2)
	fmt.Printf("%v\n", num1^num2)
	fmt.Printf("%v\n", num1&^num2)
}

/*SUMMARY
1. Variable Declaration
	-	Variables can be declared in any of the following ways in Go
		|-----------------------------------------------------------------------
		Example 1: var a int
					a = 10
		Example 2: var a int = 10
		Example 3: var a := 10
	- Variables can't be redeclared but can be shadowed (Shadowed here means when globally declared, it can be assigned a different value in local scope)
	- All declared variables must be used

2. Variable Visibility
	- lower case first letter for package scope
	- uppercase first letter to export
	- no private scope in Golang

3. Naming conventions
	- Pascal or Camel case
	- Capitalize acronyms e.g I prefer thisURL, thisHTTP
4. Type Conversions
	- use strconv for string conversion https://pkg.go.dev/search?limit=25&m=package&q=strconv
*/
