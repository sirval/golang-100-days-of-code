package main

import (
	"fmt"
)

const conferenceTicket uint = 50

var conferenceName string = "Go Conference"
var remainingTicket uint
var bookings = []string{}

func main() {
	greet(conferenceTicket, remainingTicket, conferenceName)
	for {
		firstName, lastName, email, country, userTickets := getUserInputs()
		remainingTicket = conferenceTicket - userTickets

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookings = append(bookings, firstName+" "+lastName)
			firstNames := getFirstName(bookings)

			fmt.Printf("First names in the booking list are: %v\n", firstNames)
			notifyByCountry(country, firstName, lastName, conferenceName)
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

func notifyByCountry(country string, firstName string, lastName string, confName string) {
	switch country {
	case "Nigeria":
		fmt.Printf("We're thrilled to have you %v %v from %v in the 2023 %v\n", firstName, lastName, country, confName)
	case "Ghana":
		fmt.Printf("We're thrilled to have you %v %v from %v in the 2023 %v\n", firstName, lastName, country, confName)
	case "Israel", "USA":
		fmt.Printf("We're thrilled to have you %v %v from %v in the 2023 %v\n", firstName, lastName, country, confName)
	default:
		fmt.Println("No valid country selected")
	}
}
