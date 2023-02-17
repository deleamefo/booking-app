package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 23
	bookings := []string{}

	fmt.Printf("Welome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//var bookings = [50]string
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		// use this (remainingTickets - uint(userTickets) =..... or place uint on userTickets as 'var userTickets uint')

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets reamining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings { // ( _ = index)
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			// noTicketsRemaining := remainingTickets == 0 // Do this when the handling more than one if stmt
			// if noTicketsRemaining { // no tickets remaining
			// 	// end program
			// 	fmt.Println("We have no tickets remaining. Come back next year!")
			// 	break

			if remainingTickets == 0 { // no tickets remaining
				// end program
				fmt.Println("We have no tickets remaining, please come next year!")
				break
			}

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			//continue
		}

	}

}
