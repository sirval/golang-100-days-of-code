package main

import "fmt"

func getUserInputs() (string, string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var country string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter your country:")
	fmt.Scan(&country)
	fmt.Println("How many tickets would you like to buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, country, userTickets
}
