package main

import "fmt"

func main() {
	//	var confrenceName = "Go Confrence"
	confrenceName := "Go Confrence" // shorthand declaration
	const confrenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to our", confrenceName, "booking application")
	fmt.Println("We have", remainingTickets, "tickets remaining out of a total of", confrenceTickets)
	fmt.Println("Get your tickets now")

	var firstName string
	var lastName string
	var userTickets int
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName) //user input pointer

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName) //user input pointer

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets) //user input pointer

	fmt.Printf("Thank you %s %s for booking %d tickets for %s\n", firstName, lastName, userTickets, confrenceName)
}
