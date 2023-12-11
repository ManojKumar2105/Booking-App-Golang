package main

import (
	"Booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var appName string = "Go conference" //Special feature available to var types alone where we can't declare datatypes explicitly and same as var
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0) // List of maps so that empty block braces are present inside the function
var bookings = make([]UserData, 0) // List of maps so that empty block braces are present inside the function
const totalTickets uint = 50

type UserData struct {
	firstName     string
	lastName      string
	email         string
	ticketsNeeded uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, ticketsNeeded := getUserInput()

	isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, ticketsNeeded, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {

		ticketBooking(firstName, lastName, ticketsNeeded, email)

		var noOfTickets bool = remainingTickets == 0
		if noOfTickets {
			fmt.Println("All Tickets were sold out...Come back next year")

		}

	} else {

		if !isValidName {
			fmt.Printf("Check the first name and last name you have entered\n")
		}
		if !isValidEmail {
			fmt.Printf("Check the email you have entered\n")
		}
		if !isValidTickets {
			fmt.Printf("Sorry sir, we're having %v tickets remaining\n", remainingTickets)
		}

		// continue Here continue is not needed
	}

}

func greetUsers() {
	fmt.Printf("Welcome to our %v application\n", appName)
	fmt.Printf("We have %v tickets and %v seats are available\n", totalTickets, remainingTickets)
	fmt.Println("Book your tickets and enjoy the day")
}

func PrintFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// fmt.Println(names)
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var ticketsNeeded uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter the no of tickets you needed:")
	fmt.Scan(&ticketsNeeded)

	return firstName, lastName, email, ticketsNeeded
}

func ticketBooking(firstName string, lastName string, ticketsNeeded uint, email string) {

	var userData = UserData{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		ticketsNeeded: ticketsNeeded,
	}

	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - ticketsNeeded

	wg.Add(1)
	go sendTicket(ticketsNeeded)

	firstNames := PrintFirstNames()

	fmt.Printf("List of Bookings : %v", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive the confirmation and payment mode to your registered mail:%v\n", firstName, lastName, ticketsNeeded, email)
	fmt.Printf("There are still %v seats available\n", remainingTickets)
	fmt.Printf("These are all our bookings : %v\n", bookings)
	fmt.Printf("These are all firstNames of our bookings : %s\n", firstNames)
	wg.Wait()
}

func sendTicket(ticketsNeeded uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("You have booked %v tickets from %v tickets", ticketsNeeded, remainingTickets)
	fmt.Println("################################")
	fmt.Println(ticket)
	fmt.Println("################################")
	wg.Done()
}
