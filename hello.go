package main

import ("fmt"
        "strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
    bookings := []string{}

	fmt.Println("conference")
	fmt.Printf("welcome to our %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still avalible \n", conferenceTickets, remainingTickets)
	fmt.Println("get yout tickets here to attend")

    for{
        var firstName string
        var lastName string
        var email string
        var userTicket uint

        // scan para tomar el valor asignado del input
        fmt.Println("enter your frist name: ")
        fmt.Scan(&firstName)

        fmt.Println("enter your last name: ")
        fmt.Scan(&lastName)

        fmt.Println("enter you email: ")
        fmt.Scan(&email)

        fmt.Println("enter number of tickets: ")
        fmt.Scan(&userTicket)

        remainingTickets = remainingTickets - userTicket
        bookings = append(bookings, firstName + " " + lastName)
        
        //imprimir
        fmt.Printf("Thank you %v %v Booked %v tickets. You will receive a confirmation email at %v \n ", firstName, lastName, userTicket, email)
        fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
        
        //para recorrer la lista array  solo mostrando el primer nombre del usuario ingresado
        //se importo "string" field
        firstNames := []string{}
        for _, booking := range bookings{
            var names = strings.Fields(booking)
            firstNames = append(firstNames,names[0])
        }

        // fmt.Printf("these are all our bookings: %v\n",bookings)
        fmt.Printf("the first name of bookins are: %v\n",firstNames)
    }
	
    

	

    //array para visualizar los tipos de datos, que almacena , y agrega
    // fmt.Printf("the whole array: %v\n",bookings)
    // fmt.Printf("the first value: %v\n",bookings[0])
    // fmt.Printf("array type: %T\n ",bookings)
    // fmt.Printf("array length %v\n",len(bookings))

    // fmt.Printf("Slice type: %T\n ",bookings)
    // fmt.Printf("Slice length: %v\n ",len(bookings))

}
