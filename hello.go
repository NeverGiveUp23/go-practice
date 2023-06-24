package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to our booking app")
	fmt.Println("Enter your full name")

	getCustomerName("Felix", "Vargas")

}

func getCustomerName(name string, lastName string) {
	fmt.Scan(&name, &lastName)
	str1 := strings.ToLower(name)
	str2 := strings.ToLower(lastName)

	if str1 == "felix" && str2 == "vargas" {
		useCustomerTicket()
	} else {
		fmt.Println("you have no ticket")
	}
}

func useCustomerTicket() {
	companyTickets := 10
	var customerTicket int

	fmt.Println("We have ", companyTickets, " tickets remaining... How many would you like to buy?")
	fmt.Scanln(&customerTicket)

	if customerTicket <= 0 || customerTicket > 10 {
		fmt.Println("Sorry, you need to enter a ticket amount above 0 and less than 10")
		useCustomerTicket()
	} else {
		companyTicketsLeft := companyTickets - customerTicket
		fmt.Println("You now have ", customerTicket, " and we have ", companyTicketsLeft, ".")
	}

}
