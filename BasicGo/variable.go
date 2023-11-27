package main

import "fmt"

func main() {
	// var name string
	// var name = "Abhijeet"

	name := "Abhijeet" //auto type declaration
	fmt.Println(name)

	name = "Abhijeet"
	fmt.Println(name)

	var (
		firstName  = "Abhijeet"
		middleName = "Arun"
		lastName   = "Dhumal"
	)

	fmt.Println(firstName, middleName, lastName)
}

/* >>
Abhijeet
Abhijeet
Abhijeet Arun Dhumal
*/
