package helper

import "strings"


///function to validate user inputs
func ValidateUserInput(firstName string,lastName string,email string, userTickets uint,remainigTickets uint) (bool,bool,bool) {

			isValidName  := len(firstName)>=2 && len(lastName)>=2
			isValidEmail :=strings.Contains(email,"@")
			isValidTicketNumber :=userTickets > 0 && userTickets <= remainigTickets
	return isValidName,isValidEmail,isValidTicketNumber
}