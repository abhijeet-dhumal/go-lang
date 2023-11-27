package main

import "fmt"

func main() {
	var confernceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to '", confernceName, "' conference!")
	fmt.Printf("We have total of %v tickets and %v are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend conference..")
}
