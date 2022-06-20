package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperationWithoutAssetPrice(t *testing.T) {
	ope := NewOperation(
		10.0,
		"2022-06-19",
		0.1,
		"BTC",
		0,
		PURCHASE,
		"test",
	)

	assert.Equal(t, 100.0, ope.GetAssetPrice())
}

func TestNewOperationWithAssetPrice(t *testing.T) {
		ope := NewOperation(
		10.0,
		"2022-06-19",
		0.1,
		"BTC",
		200,
		PURCHASE,
		"test",
	)

	assert.Equal(t, 200.0, ope.GetAssetPrice())
}