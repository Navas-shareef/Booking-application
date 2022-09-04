package main

import (
	"fmt"
	"strings"
)


func  main()  {

	 conferenceName  := "Go conference"
	const conferenceTickets int=50 
	var remainigTickets uint =50 
	
	bookings:=  []string{}
	



	fmt.Printf("conferenceTickets is %T, conferenceName is %T,remainingTicket is %T",conferenceTickets,conferenceName,remainigTickets)

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v and  %v are still available\n",conferenceTickets,remainigTickets)
	fmt.Println("Get your tickets here to attend")
	

	for		remainigTickets >0 && len(bookings) < 50	{	
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

			isValidName  := len(firstName)>=2 && len(lastName)>=2
			isValidEmail :=strings.Contains(email,"@")
			isValidTicketNumber :=userTickets > 0 && userTickets <= remainigTickets
			
			if isValidName && isValidEmail && isValidTicketNumber{

				remainigTickets=remainigTickets - userTickets
				bookings=append(bookings,firstName+" "+lastName)
				
				fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation emait at %v\n",firstName,lastName,userTickets,email)
				fmt.Printf("%v tickets remaining for %v\n",remainigTickets,conferenceName)
				firstNames:=[]string{}			
				for _,booking := range bookings{

					var names=	strings.Fields(booking)
					firstNames=append(firstNames,names[0])
				}

				fmt.Printf("The first names of bookings are: %v\n",firstNames)

			}else if userTickets==remainigTickets{

					//end program
					fmt.Println("Our conference is booked out. Come back next year. ")
					break
				
			}else{
				if !isValidName{
					fmt.Println(" first name or last name you entered is too short")
				}
				if !isValidEmail{
					fmt.Println("email address you entered doesn't contain @ sign")
				}
				if !isValidTicketNumber{
					fmt.Println("invalid userToken")
				}

			
			}
			
		
			
			
		
	}	
	
	
	
}

