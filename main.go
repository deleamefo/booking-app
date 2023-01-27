package main

import "fmt"

func main() {

	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 23

	fmt.Printf("Welome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name

	fmt.Println("Enter your first name:")
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
