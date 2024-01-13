package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const confrenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, confrenceName is %T\n", confrenceTickets, remainingTickets, conferenceName)

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
