package main

import (
	"fmt"
	"parking-lot-2/parking"
	"parking-lot-2/subscription"
)

func main() {

	var input int

	fmt.Println("Welcome to parking apps!\nPlease choose available menu below:")
	fmt.Println("1. Subscribe\n2. Parking\n3. Unparking\n4. Exit")
	fmt.Scanln(&input)

	if input == 1 {
		subs := subscription.GetSubsID()
		fmt.Print(subs)

	} else if input == 2 {

		var inputCar string
		fmt.Println("Please input your car plate number: ")
		fmt.Scanln(&inputCar)

		var plate parking.Parking
		plate.Plate = inputCar
		plate.GetPlateNumber(plate.Plate)

		ticket := "Your ticket ID is " + parking.GetTicketID()
		fmt.Println(ticket)

	} else if input == 3 {
		var inputCar string
		fmt.Println("Please input your car plate number: ")
		fmt.Scanln(&inputCar)

		var plate parking.Parking
		plate.Plate = inputCar
		plate.CarUnparking(plate.Plate)

	} else if input == 4 {
		fmt.Println("Thank you")
	}
}
