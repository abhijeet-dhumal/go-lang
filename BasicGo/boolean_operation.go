package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	// && = and
	fmt.Println(a && b) // false
	// || = or
	fmt.Println(a || b) // true
	// ! = not
	fmt.Println(!a) // false

	// Boolean operation program
	var var1 = 90
	var var2 = 80

	var myvar1 bool = var1 > 80
	var myvar2 bool = var2 > 80

	var myVarsAnd bool = myvar1 && myvar2
	var myVarsOr bool = myvar1 || myvar2

	fmt.Println("MyVars And : Or op - '", myVarsAnd, ":", myVarsOr, "'")
}

/* >>
false
true
false
MyVars And : Or op - ' false : true '
*/
