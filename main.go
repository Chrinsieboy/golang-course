package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const confrenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, confrenceName is %T\n", confrenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for name

	userName = "Chris"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
}
