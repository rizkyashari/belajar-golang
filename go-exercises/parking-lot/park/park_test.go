package park_test

import (
	"parking-lot/car"
	"parking-lot/park"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarParkingNormal(t *testing.T) {
	park := park.Park{
		CustomerCar:  map[string]car.Car{},
		CarPlate:     map[string]string{},
		AvailableLot: 10,
	}
	car1 := car.Car{Plate: "KT1234AB"}
	car1ParkingTicket, err := park.CarParking(car1)
	assert.NotEmpty(t, car1ParkingTicket)
	assert.Nil(t, err)
	assert.Equal(t, 9, park.AvailableLot)
	assert.Equal(t, 1, len(park.CustomerCar))
	assert.Equal(t, 1, len(park.CarPlate))
}

func CarUnparkingWrongTicket(t *testing.T) {
	park := park.Park{
		CustomerCar:  map[string]car.Car{},
		CarPlate:     map[string]string{},
		AvailableLot: 1,
	}

	car1 := car.Car{Plate: "KT1234AB"}
	park.CarParking(car1)
	c, err := park.CarUnparking("wrong ticket")
	assert.NotEqual(t, car1, c)
	assert.NotNil(t, err)
	assert.Equal(t, 0, park.AvailableLot)
	assert.Equal(t, 1, len(park.CustomerCar))
	assert.Equal(t, 1, len(park.CarPlate))
}

func TestString(t *testing.T) {
	park := park.Park{
		CustomerCar:  map[string]car.Car{},
		CarPlate:     map[string]string{},
		AvailableLot: 1,
	}
	assert.Equal(t, "Lot: map[]\nAvailableLot: 1", park.String())
}
