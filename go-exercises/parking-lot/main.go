package main

import (
	"fmt"
	"parking-lot/attendant"
	"parking-lot/car"
	"parking-lot/park"
)

func main() {
	park1 := park.Park{
		CustomerCar:  map[string]car.Car{},
		CarPlate:     map[string]string{},
		AvailableLot: 2,
	}

	attendant1 := attendant.Attendant{
		ID:   0,
		Park: &park1,
	}
	for {
		var input int
		fmt.Println("Choose the given menu below: ")
		fmt.Println("1. Park\n2. Unpark\n3. Exit")
		fmt.Scanln(&input)

		if input == 1 {
			carPark(&attendant1)
		} else if input == 2 {
			carUnpark(&attendant1)
		} else if input == 3 {
			fmt.Println("Thank you.")
			break
		} else {
			fmt.Println("Invalid input.")
		}
	}
}

func carPark(atd *attendant.Attendant) {
	fmt.Println("==PARK==")
	fmt.Print("Enter your car plate: ")
	var plate string
	fmt.Scanln(&plate)
	custCar := car.Car{Plate: plate}
	ticket, err := atd.CarParking(custCar)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Park car with plat", plate, "your ticket is", ticket)
}

func carUnpark(atd *attendant.Attendant) {
	fmt.Println("==UNPARK==")
	fmt.Print("Enter your ticket: ")
	var ticket string
	fmt.Scanln(&ticket)
	custCar, err := atd.CarUnparking(ticket)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Unpark car with plat", custCar.Plate)
}
