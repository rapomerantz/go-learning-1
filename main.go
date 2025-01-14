package main

import (
	"fmt"
	"strings"
)

const confrenceTickets = 50

var confrenceName = "Go Confrence" // shorthand declaration invalid for package level variables
var remainingTickets = 50
var bookings []string //this is a slice, similar to array but with dynamic size

func main() {

	greetUsers()

	for { //infinite loop
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTickets {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets have been sold out")
				break
			}

		} else {
			fmt.Println("invalid input")
		}
	}
}

func greetUsers() {
	fmt.Println("!!!Welcome to our", confrenceName, "booking application")
	fmt.Println("We have", remainingTickets, "tickets remaining out of a total of", confrenceTickets)
	fmt.Println("Get your tickets now")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		//_ is a throwaway variable (blank identifier)
		var names = strings.Fields(booking) //creates an array of strings from the booking string
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var userTickets int
	var email string

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName) //user input pointer

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName) //append to the slice. have to assign back to variable

	fmt.Printf("Thank you %s %s for booking %d tickets for %s, we will email you at %s\n", firstName, lastName, userTickets, confrenceName, email)
	fmt.Printf("We have %d tickets remaining out of a total of %d\n", remainingTickets, confrenceTickets)
}
