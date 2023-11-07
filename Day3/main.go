package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket uint = 50
	var remainingTicket uint
	bookings := []string{}

	fmt.Printf("conferenceTicket is %T, remainingTicket is %T, conferenceName is %T\n", conferenceTicket, remainingTicket, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")

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
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("User %v %v booked %v tickets for the %v.\nWe'll send you the necessary details via your email at %v\n", firstName, lastName, userTickets, conferenceName, email)
	// fmt.Printf("We have %v available tickets. Feel free to let your friends know about the %v\n", remainingTicket, conferenceName)
	fmt.Printf("The whole array type: %v\n", bookings)
	fmt.Printf("The first value in the array: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("All bookings %v\n", bookings)
}
