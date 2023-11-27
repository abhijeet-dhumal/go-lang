package main

func main() {
	// const firstName string = "Abhijeet"
	// const lastName = "Dhumal"

	const (
		firstName string = "Abhijeet"
		lastName         = "Dhumal"
	)
	firstName = "Dhumal" //>> Error : ./constant.go:11:2: cannot assign to firstName (constant "Abhijeet" of type string)
	// error
	// firstName = "Dhumal"
	// lastName = "Abhijeet"
}
