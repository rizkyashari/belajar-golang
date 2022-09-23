package lots_test

import (
	"parking-lot/car"
	"parking-lot/lots"
	"parking-lot/park"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAvailableLots(t *testing.T) {
	lots1 := lots.Lots{
		ParkingLots: []*park.Park{
			{
				CustomerCar:  map[string]car.Car{},
				CarPlate:     map[string]string{},
				AvailableLot: 0,
			},
			{
				CustomerCar:  map[string]car.Car{},
				CarPlate:     map[string]string{},
				AvailableLot: 7,
			},
			{
				CustomerCar:  map[string]car.Car{},
				CarPlate:     map[string]string{},
				AvailableLot: 9,
			},
		},
	}
	park, err := lots1.GetAvailableLots()
	assert.Equal(t, 7, park.AvailableLot)
	assert.Nil(t, err)
}
