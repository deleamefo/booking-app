package main

import (
	"fmt"
	"strconv"
)

const conferenceTickets int = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) // bookings := []string{}

func main() {

	greetUsers()
	for {
		// function to get user info
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)
			// call function print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			if remainingTickets == 0 { // no tickets remaining
				// end program
				fmt.Println("We have no tickets remaining, please come next year!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short, pls try again.")
			}
			if !isValidEmail {
				fmt.Println("Invalid email")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

		}

	}

}

// function calls
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // ( '_' = index )

		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {

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

	return firstName, lastName, email, userTickets
}

// function to book tickets
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets reamining for %v\n", remainingTickets, conferenceName)
}
