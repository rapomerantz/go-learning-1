package main

import (
	"fmt"
	"strings"
)

func main() {
	//	var confrenceName = "Go Confrence"
	confrenceName := "Go Confrence" // shorthand declaration
	const confrenceTickets = 50
	var remainingTickets = 50
	var bookings []string //this is a slice, similar to array but with dynamic size

	greetUsers(confrenceName, confrenceTickets, remainingTickets)

	for {
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

		isValidName := len(lastName) >= 2 && len(firstName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName) //append to the slice. have to assign back to variable

			fmt.Printf("Thank you %s %s for booking %d tickets for %s\n", firstName, lastName, userTickets, confrenceName)
			fmt.Printf("We have %d tickets remaining out of a total of %d\n", remainingTickets, confrenceTickets)

			firstNames := []string{}

			for _, booking := range bookings {
				//_ is a throwaway variable (blank identifier)
				var names = strings.Fields(booking) //creates an array of strings from the booking string
				firstNames = append(firstNames, names[0])
			}

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

func greetUsers(confrenceName string, remainingTickets int, confrenceTickets int) {
	fmt.Println("!!!Welcome to our", confrenceName, "booking application")
	fmt.Println("We have", remainingTickets, "tickets remaining out of a total of", confrenceTickets)
	fmt.Println("Get your tickets now")
}
