package main

import (
	"booking-app/helper"
	"fmt"	
	"strconv"
)


	var conferenceName  = "Go conference"
	const conferenceTickets int=50 
	var remainigTickets uint =50 
	var bookings =  make([]map[string]string,0)


func  main()  {


	
	greetUsers()

	for		remainigTickets >0 && len(bookings) < 50	{	
			
			firstName,lastName,email,userTickets := getUserInput()
			isValidName,isValidEmail,isValidTicketNumber :=helper.ValidateUserInput(firstName,lastName,email,userTickets,remainigTickets)

			if isValidName && isValidEmail && isValidTicketNumber{

				bookTicket(userTickets,firstName,lastName,email,conferenceName)

				firstNames:=getFirstNames()
				fmt.Printf("The first names of bookings are: %v\n",firstNames)

			}else if remainigTickets==0{

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



/// wishing user
func  greetUsers()  {
	fmt.Printf("Welcome to %v bookong application\n",conferenceName)
	fmt.Printf("We have total of %v and  %v are still available\n",conferenceTickets,remainigTickets)
	fmt.Println("Get your tickets here to attend")
}


/// displaying user first names
func getFirstNames()[]string  {
	firstNames:=[]string{}			
	for _,booking := range bookings{
		// var names=	strings.Fields(booking)
		firstNames=append(firstNames,booking["firstName"])
	}
	return firstNames
}


///reading user inputs
func getUserInput()  (string,string,string,uint){
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

		return firstName,lastName,email,userTickets	
}


func bookTicket(userTickets uint ,firstName string,lastName string ,email string,conferenceName string)  {
	remainigTickets=remainigTickets - userTickets

	//create a map for a user
	var userData=make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets),10)




	bookings=append(bookings,userData)

	fmt.Printf("List of bookings is %v\n",bookings)
				
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation emait at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remainigTickets,conferenceName)
					
}