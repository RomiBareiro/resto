package server

import (
	"errors"
	"testing"

	"resto_go/types"

	"github.com/stretchr/testify/assert"
)

func TestValidateInputData(t *testing.T) {
	validInput, err := ValidateInputData("40.7128", "-74.0060")
	assert.NoError(t, err, "No error")
	expectedValidInput := types.InputData{Latitude: 40.7128, Longitude: -74.0060}
	assert.Equal(t, expectedValidInput, validInput, "TEST ok")

	_, err = ValidateInputData("invalid", "-74.0060")
	assert.Error(t, err, "Error was expected")
	assert.True(t, errors.Is(err, errInvalidParams), "Expected invalid params error")

	_, err = ValidateInputData("0", "0")
	assert.Error(t, err, "Expected error missing fields, test failed")
}
