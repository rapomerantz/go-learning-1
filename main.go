package main

import (
	"booking-app/common" //importing a package from a different directory references module name
	"fmt"
	"time"
)

const confrenceTickets = 50

var confrenceName = "Go Confrence" // shorthand declaration invalid for package level variables
var remainingTickets = 50

// var bookings = make([]map[string]string, 0) //this is a slice, similar to array but with dynamic size. because we're using map, we need to use make to initialize it
var bookings = make([]UserData, 0) //replacing with a slices of struct

// struct allows mixed types, a lightweight class
// type keyword creates a custom type in the app
type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   int
}

func main() {

	greetUsers()

	for { //infinite loop
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := common.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTickets {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, email) // go keyword abstracts away conconcurent threading

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
		//booking variable is a map
		firstNames = append(firstNames, booking.firstName)
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

	var userData = UserData{
		firstName,
		lastName,
		email,
		userTickets,
	} //creating a struct

	bookings = append(bookings, userData) //append to the slice. have to assign back to variable
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %s %s for booking %d tickets for %s, we will email you at %s\n", firstName, lastName, userTickets, confrenceName, email)
	fmt.Printf("We have %d tickets remaining out of a total of %d\n", remainingTickets, confrenceTickets)
}

func sendTicket(userTickets int, firstName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets to %v user via email %v\n", userTickets, firstName, email) //sprintf formats and returns a string
	fmt.Println("#####")
	fmt.Printf("sending ticket: \n%v to email address %v\n", ticket, email)
	fmt.Println("#####")
}
