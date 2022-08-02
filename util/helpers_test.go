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

func TestTransformStringSliceIntoInterfaceHappyPath(t *testing.T) {
	result := TransformStringSliceIntoInterface([]string{"Bonjour", "hola"})

	assert.Equal(t, []interface{}{"Bonjour", "hola"}, result)
}

func TestTransformStringToFloatWithEmptyString(t *testing.T) {
	assert.Equal(t, float64(0), TransformStringToFloat(""))
}

func TestTransformStringToFloatWithHappy(t *testing.T) {
	assert.Equal(t, float64(24.2), TransformStringToFloat("24.2"))
	assert.Equal(t, float64(10.0), TransformStringToFloat("10"))
	assert.Equal(t, float64(1.84245), TransformStringToFloat("1.84245"))
}

// func TestTransformStringToFloatWithErrorWhenConvertingFloat(t *testing.T) {
// 	// val := TransformStringToFloat("qohfoqhofqhfqiohfhiofiohqohqiiofqhqioh")

// 	exiter = New(TransformStringToFloat("qohfoqhofqhfqiohfhiofiohqohqiiofqhqioh"))

// 	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
// 		return
// 	}

// 	assert.Fail(
// 		t,
// 		"Error parsing Float:  strconv.ParseFloat: parsing \"qohfoqhofqhfqiohfhiofiohqohqiiofqhqioh\": invalid syntax",
// 		TransformStringToFloat("qohfoqhofqhfqiohfhiofiohqohqiiofqhqioh"),
// 	)
// }

func TestReverSliceHappyPath(t *testing.T) {
	result := ReverseSlice([]string{"test", "word", "not", "in"})

	assert.Equal(t, []string{"in", "not", "word", "test"}, result)
}

func TestRemoveKeyFromSlice(t *testing.T) {
	result := RemoveKeyFromSlice([]string{"test", "super", "toto"}, 1)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, []string{"test", "toto"}, result)
}

func TestRemoveKeyFromSliceLastKey(t *testing.T) {
	result := RemoveKeyFromSlice([]string{"test", "super", "toto"}, 2)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, []string{"test", "super"}, result)
}
