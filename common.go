package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(lastName) >= 2 && len(firstName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets //go can return multiple values
}
