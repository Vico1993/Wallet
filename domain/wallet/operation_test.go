package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithoutUnitPrice(t *testing.T) {
	operation := NewOperation(
		"2022-06-19",
		0.1,
		"BTC",
		0,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	)

	assert.Equal(t, float64(100), operation.UnitPrice)
}