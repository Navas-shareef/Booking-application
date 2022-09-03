package main

import "fmt"

func  main()  {

	var conferenceName  = "Go conference"
	const conferenceTickets=50 
	var remainigTickets=50 

	fmt.Printf("conferenceTickets is %T, conferenceName is %T,remainingTicket is %T",conferenceTickets,conferenceName,remainigTickets)

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v and  %v are still available\n",conferenceTickets,remainigTickets)
	fmt.Println("Get your tickets here to attend")
	

	var userName string
	var userTickets int
	// ask user for their name

	userName="Tom"
	userTickets=2
	fmt.Printf("User %v booked %v tickets\n",userName,userTickets)


}

