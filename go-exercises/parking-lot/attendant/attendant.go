package attendant

import (
	"parking-lot/car"
	"parking-lot/park"
)

type Attendant struct {
	ID   int
	Park *park.Park
}

func (a *Attendant) CarParking(c car.Car) (string, error) {
	ticket, err := a.Park.CarParking(c)
	if err != nil {
		return "", err
	}
	return ticket, nil
}

func (a *Attendant) CarUnparking(ticket string) (car.Car, error) {
	custCar, err := a.Park.CarUnparking(ticket)
	if err != nil {
		return custCar, err
	}
	return custCar, nil
}
