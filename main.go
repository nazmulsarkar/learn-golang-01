package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Booking App for the conference!")
	fmt.Println("Book your ticket(s) before booking close!")

	const totalTickets = 50
	fName := ""
	lName := ""
	email := ""
	var remainingTickets uint = 50
	var bookedTickets uint = 0

	var bookings []string = []string{}

	for remainingTickets > 0 && len(bookings) < totalTickets {

		fmt.Println("")
		fmt.Println("Enter first name:")
		fmt.Scan(&fName)
		fmt.Println("-----------------------")

		fmt.Println("")
		fmt.Println("Enter last name:")
		fmt.Scan(&lName)
		fmt.Println("-----------------------")

		fmt.Println("")
		fmt.Println("Enter email:")
		fmt.Scan(&email)
		fmt.Println("-----------------------")

		fmt.Println("")
		fmt.Println("How many tickets you require?")
		fmt.Scan(&bookedTickets)
		fmt.Println("-----------------------")

		if bookedTickets > remainingTickets {
			fmt.Println("You can't book more than available tickets.")
			continue
		}

		name := fName + " " + lName

		remainingTickets = remainingTickets - bookedTickets

		fmt.Printf("Thank you %v for booking %v tickets.\n", name, bookedTickets)
		fmt.Printf("There'll be a confirmation email at %v.", email)
		fmt.Print("\n")
		fmt.Printf("There're only %v tickets remaining from %v.\n", remainingTickets, totalTickets)
		fmt.Println("")

		bookings = append(bookings, name)

		fmt.Print(bookings)
	}

}
