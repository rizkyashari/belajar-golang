package parking

import (
	"fmt"
	"math/rand"
)

type Parking struct {
	Plate        string
	Ticket       string
	PlateNumber  map[int]string
	TicketID     map[int]string
	AvailableLot int
}

func (parking *Parking) GetPlateNumber(plate string) {
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

func (parking *Parking) CarUnparking(plate string) {
	fmt.Println("Your car with plate number", parking.Plate, "has been unparked")
}
