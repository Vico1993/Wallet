package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperationWithoutUnitPrice(t *testing.T) {
	ope := NewOperation(
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

	assert.Equal(t, 100.0, ope.GetUnitPrice())
}

func TestNewOperationWithUnitPrice(t *testing.T) {
	ope := NewOperation(
		"2022-06-19",
		0.1,
		"BTC",
		200,
		"fiat",
		10.0,
		10.0,
		"CAD",
		PURCHASE,
		"test",
	)

	assert.Equal(t, 200.0, ope.GetUnitPrice())
}