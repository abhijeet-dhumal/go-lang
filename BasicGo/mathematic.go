package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d*e
	fmt.Println("Result c =", c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println("Result i =", i)
	i -= 5 // i = i - 5
	fmt.Println("Result i =", i)

	var q = 1
	q++ // q = q + 1
	fmt.Println("Result q =", q)
	q-- // q = q - 1
	fmt.Println("Result q =", q)
}
