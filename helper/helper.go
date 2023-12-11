package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, ticketsNeeded uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := ticketsNeeded > 0 && ticketsNeeded <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}
