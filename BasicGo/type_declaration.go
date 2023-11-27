package main

import "fmt"

func main() {
	type stringType string
	var stringVar stringType = "1234567890"

	var stringVar2 string = "0987654321"
	var stringVar2Ktp stringType = stringType(stringVar2)

	fmt.Println(stringVar)
	fmt.Println(stringVar2)
	fmt.Println(stringVar2Ktp)
}

/*>>
1234567890
0987654321
0987654321
*/
