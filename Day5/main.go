package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket uint = 50
	var remainingTicket uint
	bookings := []string{}

	fmt.Printf("conferenceTicket is %T, remainingTicket is %T, conferenceName is %T\n", conferenceTicket, remainingTicket, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("How many tickets would you like to buy:")
		fmt.Scan(&userTickets)
		remainingTicket = conferenceTicket - userTickets

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 &&
			userTickets <= conferenceTicket &&
			userTickets <= remainingTicket

		if isValidName && isValidEmail && isValidTicketNumber {
			bookings = append(bookings, firstName+" "+lastName)

			firstNames := []string{}           //Slice
			for _, booking := range bookings { //underscore tells go to overlook a variable
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("First names in the booking list are: %v\n", firstNames)

		} else {
			if !isValidName {
				fmt.Println("First name or last name must be greater than 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Email must contain @ character")
			}
			if !isValidTicketNumber {
				fmt.Printf("Ticket number must not be greater than %v\n", remainingTicket)
			}
		}
	}
}
