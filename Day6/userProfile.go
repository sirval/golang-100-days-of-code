package main

import "strings"

func getFirstName(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings { //underscore tells go to overlook a variable
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
