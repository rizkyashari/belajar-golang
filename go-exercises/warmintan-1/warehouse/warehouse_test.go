package warehouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStockFromChef(t *testing.T) {
	v := GetStockFromChef("Mie")
	assert.Equal(t, v, map[string]string{"": "Mie"})
}
