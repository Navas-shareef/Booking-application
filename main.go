package main

import "fmt"

func  main()  {

	var conferenceName  = "Go conference"
	const conferenceTickets=50 
	var remainigTickets uint =50 
	
	



	fmt.Printf("conferenceTickets is %T, conferenceName is %T,remainingTicket is %T",conferenceTickets,conferenceName,remainigTickets)

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v and  %v are still available\n",conferenceTickets,remainigTickets)
	fmt.Println("Get your tickets here to attend")
	

	var firstName string
	var lastName string
	var email string 
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)


	remainigTickets=remainigTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation emait at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remainigTickets,conferenceName)

}

