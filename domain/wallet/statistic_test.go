package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculProfitWithNoGain(t *testing.T) {
	result := calculProfit(50, 50)

	assert.Equal(t, float64(0), result, "Result should be equal to 0")
}

func TestCalculProfitGain(t *testing.T) {
	result := calculProfit(50, 100)

	assert.Equal(t, float64(100), result, "Result should be equal to 100")
}

func TestCalculProfitWithNegativeGain(t *testing.T) {
	result := calculProfit(50, 10)

	assert.Equal(t, float64(-80), result, "Result should be equal to -80")
}

func TestCalculProfitWithLargeGain(t *testing.T) {
	result := calculProfit(1, 10)

	assert.Equal(t, float64(900), result, "Result should be equal to 900")
}