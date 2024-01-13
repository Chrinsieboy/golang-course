package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const confrenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("We have a total of", confrenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for name

	userName = "Chris"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
}
