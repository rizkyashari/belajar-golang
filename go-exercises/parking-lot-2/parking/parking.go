package parking

import (
	"errors"
	"fmt"
	"math/rand"
)

type Parking struct {
	Plate        string
	Lot          map[string]string
	AvailableLot int
}

func (parking Parking) GetPlateNumber(plate string) {
	fmt.Println("Your car plate number is", parking.Plate)
}

func GetTicketID() string {
	Letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	a := make([]rune, 6)
	for i := range a {
		a[i] = Letters[rand.Intn(len(Letters))]
	}
	return string(a)
}

func (parking *Parking) CarParking(plate string) (string, error) {
	if parking.AvailableLot <= 0 {
		return "", errors.New("Lot unavailable")
	}
	if _, ok := parking.Lot[plate]; ok {
		return "", errors.New("You can't park twice")
	}
	ticket := GetTicketID()
	parking.Lot[plate] = ticket
	parking.AvailableLot -= 1

	return ticket, nil
}

func (parking *Parking) CarUnparking(plate string) {
	fmt.Println("Your car with plate number", parking.Plate, "and ticket ID", GetTicketID(), "has been unparked")
}
