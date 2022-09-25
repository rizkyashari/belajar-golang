package main

import (
	"testing"
	"warmintan-1/chef"
	"warmintan-1/waiter"
	"warmintan-1/warehouse"

	"github.com/stretchr/testify/assert"
)

func TestAddStock(t *testing.T) {
	v := waiter.AddStock("Mie")
	assert.Equal(t, v, map[string]string{"": "Mie"})
}

func TestGetStockFromWaiter(t *testing.T) {
	v := chef.GetStockFromWaiter("Mie")
	assert.Equal(t, v, map[string]string{"": "Mie"})
}
func TestGetStockFromChef(t *testing.T) {
	v := warehouse.GetStockFromChef("Mie")
	assert.Equal(t, v, map[string]string{"": "Mie"})
}
