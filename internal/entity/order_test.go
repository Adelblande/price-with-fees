package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErroIfIdIsBlank(t *testing.T) {
	order := Order {}
	assert.Error(t, order.Validate(), "id is requered")
}

func TestIfItGetsAnErroIfPriceIsBlank(t *testing.T) {
	order := Order {ID: "123"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfItGetsAnErroIfTaxIsBlank(t *testing.T) {
	order := Order {ID: "123", Price: 10}
	assert.Error(t, order.Validate(), "tax invalid")
}

func TestIfItGetsAnErroIfPriceIsInvalid(t *testing.T) {
	order := Order {ID: "123", Price: -10, Tax: 2}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func TestIfItGetsAnErroIfTaxIsInvalid(t *testing.T) {
	order := Order {ID: "123", Price: 10, Tax: -2}
	assert.Error(t, order.Validate(), "tax invalid")
}

func TestFinalPrice(t *testing.T) {
	order := Order {ID: "123", Price: 10, Tax: 2}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 12.0, order.FinalPrice)
}