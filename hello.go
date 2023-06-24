package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Enter Your Name: ")
	getCustomerName()

}

func getCustomerName() {
	var name string
	fmt.Scanln(&name)
	str1 := strings.ToLower(name)

	if str1 != "" {
		fmt.Println("Welcome to our booking app", name)
		useCustomerTicket(name)
	} else {
		fmt.Println("enter your name")
		getCustomerName()
	}
}

func useCustomerTicket(name string) {
	companyTickets := 10
	var customerTicket int

	fmt.Println("We have ", companyTickets, " tickets remaining... How many would you like to buy?")
	fmt.Scanln(&customerTicket)

	if customerTicket <= 0 || customerTicket > 10 {
		fmt.Println("Sorry, you need to enter a ticket amount above 0 and less than 10")
		useCustomerTicket(name)
	} else {
		companyTicketsLeft := companyTickets - customerTicket
		fmt.Println("You now have ", customerTicket, " tickets", name, " and we have ", companyTicketsLeft, " tickets remaining.")
	}

}
