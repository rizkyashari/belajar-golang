package main

import (
	"fmt"
	"parking-lot-2/parking"
	"parking-lot-2/subscription"
)

func main() {

	park := parking.Parking{
		PlateNumber:  map[int]string{},
		TicketID:     map[int]string{},
		AvailableLot: 2,
	}

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

		if plate.Plate == inputCar {
			for i := 1; i < 3; i++ {
				park.PlateNumber[i] = plate.Plate
				park.TicketID[i] = parking.GetTicketID()
				park.AvailableLot -= 1

				v1 := park.PlateNumber[i]
				v2 := park.TicketID[i]
				v3 := park.AvailableLot

				fmt.Println("===PARKING===")
				fmt.Println("Plate Number:", v1)
				fmt.Println("Ticket ID:", v2)
				fmt.Println("Available Lot:", v3)
				fmt.Println("")
				fmt.Println(park)
			}
		}

	} else if input == 3 {
		var inputCar string
		fmt.Println("Please input your ticket ID: ")
		fmt.Scanln(&inputCar)

		var plate parking.Parking
		plate.CarUnparking(plate.Plate)

	} else if input == 4 {
		fmt.Println("Thank you")
	}

}
