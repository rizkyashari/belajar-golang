package waiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStock(t *testing.T) {
	v := AddStock("Mie")
	assert.Equal(t, v, map[string]string{"": "Mie"})
}

func TestTakeOrder(t *testing.T) {
	v := TakeOrder("Telur")
	assert.Equal(t, v, map[string]string{"": "Telur"})
}
