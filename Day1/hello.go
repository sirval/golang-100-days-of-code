package main

import "fmt"

//Example 1
/*|----------------------------------------------------|
|----------------------------------------------------|*/

func main() {
	fmt.Println("Welcome to Go programming")
	//variable declaration
	// var a int
	// a = 10

	//2 var a int = 12
	a := 15
	fmt.Println("c = ", a)
}

/*Variables in GoLang has 3 visibility Package level variables, scope (function) level and finally Global variables
<+-----------------------------------------------------------+>
---------------------------------------------------------------
scope or local variables will always take precedence over package visibility variables
E.G

var A string = 'hello' //globally visible ---- see uppercased variable
var a string = 'hello' //package visible ---- when declared outside the scope within a package
func main(){
	var a string = 'hello' //scope visible ---- when declared inside a function
}

Example 2*/
// var username string = "IamValentine"
// var (
// 	name    string = "Ikenna Valentine"
// 	id      int32  = 12
// 	surname        = "Ohuka"
// )

// func main() {
// 	fmt.Println(username)
// }
