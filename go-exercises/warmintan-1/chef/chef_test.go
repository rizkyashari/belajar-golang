package chef

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStockFromWaiter(t *testing.T) {
	v := GetStockFromWaiter("Mie")
	assert.Equal(t, v, map[string]string{"": "Mie"})
}

func TestGetOrderFromWaiter(t *testing.T) {
	v := GetOrderFromWaiter("Telur")
	assert.Equal(t, v, map[string]string{"": "Telur"})
}
