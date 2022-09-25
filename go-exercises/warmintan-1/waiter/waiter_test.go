package waiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStock(t *testing.T) {
	v := AddStock("Mie")
	assert.Equal(t, v, map[string]string{"": "Mie"})
}
