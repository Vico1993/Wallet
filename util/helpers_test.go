package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformFloatToString(t *testing.T) {
	result := FormatFloat(3.467)

	assert.Equal(t, "3.46", result)
}

func TestTransformFloatToStringWithVerySmallValue(t *testing.T) {
	result := FormatFloat(0.0067)

	assert.Equal(t, "0.0067", result)
}


func TestIsInStringSliceNotIn(t *testing.T) {
	result := IsInStringSlice("toto", []string{"test", "word", "not", "in"})
	assert.Equal(t, false, result)
}

func TestIsInStringSliceIn(t *testing.T) {
	result := IsInStringSlice("in", []string{"test", "word", "not", "in"})
	assert.Equal(t, true, result)
}
