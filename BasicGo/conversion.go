package main

import "fmt"

func main() {
	var var32 int32 = 32769
	var var64 int64 = int64(var32)
	var var16 int16 = int16(var32) // number overflow

	fmt.Println(var32)
	fmt.Println(var64)
	fmt.Println(var16)

	var name = "Abhijeet Dhumal"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}

/* >>
32769
32769
-32767
Abhijeet Dhumal
65
A
*/
