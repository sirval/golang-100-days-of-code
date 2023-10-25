package main

import "fmt"

func main() {
	userBill := newUserBill("John Doe Bill")
	fmt.Println(userBill.formatter())
}
