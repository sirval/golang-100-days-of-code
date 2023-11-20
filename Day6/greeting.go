package main

import "fmt"

func greet(conferenceTicket uint, remainingTicket uint, conferenceName string) {
	fmt.Printf("conferenceTicket is %T, remainingTicket is %T, conferenceName is %T\n", conferenceTicket, remainingTicket, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}
