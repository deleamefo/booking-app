package main

import "fmt"

func main() {

	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 23

	fmt.Printf("Welome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")
}
