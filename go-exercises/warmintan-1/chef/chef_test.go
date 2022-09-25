package chef

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStockFromWaiter(t *testing.T) {
	v := GetStockFromWaiter("Mie")
	assert.Equal(t, v, map[string]string{"": "Mie"})
}
