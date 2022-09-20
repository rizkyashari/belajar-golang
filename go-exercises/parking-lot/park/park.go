package park

import (
	"errors"
	"fmt"
	"math/rand"
	"parking-lot/car"
)

type Park struct {
	CustomerCar  map[string]car.Car
	AvailableLot int
	CarPlate     map[string]string
}

func (p *Park) CarParking(c car.Car) (string, error) {
	if p.AvailableLot <= 0 {
		return "", errors.New("Lot unavailable.")
	}
	if _, ok := p.CarPlate[c.Plate]; ok {
		return "", errors.New("Parking twice is not allowed.")
	}

	ticket := GenerateTicket()
	p.CustomerCar[ticket] = c
	p.CarPlate[c.Plate] = ticket
	p.AvailableLot -= 1

	return ticket, nil

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var ticketLetterLength = 3

func GenerateTicket() string {
	b := make([]rune, ticketLetterLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (p *Park) CarUnparking(ticket string) (car.Car, error) {
	custCar, ok := p.CustomerCar[ticket]
	if !ok || ticket == "" {
		return custCar, errors.New("Ticket unrecognized.")
	}

	delete(p.CustomerCar, ticket)
	delete(p.CarPlate, custCar.Plate)
	p.AvailableLot += 1

	return custCar, nil
}

func (p *Park) String() string {
	return fmt.Sprintf("CustomerCar: %v\nAvailableLot: %v", p.CustomerCar, p.AvailableLot)
}
