package server

import (
	"bytes"

	"resto_go/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInputData(t *testing.T) {
	// Testing data
	validJSON := `{"latitude": 40.7128, "longitude": -74.0060}`
	invalidJSON := `{"latitude": "invalid", "longitude": -74.0060}`
	missingFieldsJSON := `{"latitude": 0, "longitude": 0}`

	// JSON OK
	validBody := bytes.NewBufferString(validJSON)
	validInput, err := ValidateInputData(validBody)
	assert.NoError(t, err, "No error")
	expectedValidInput := types.InputData{Latitude: 40.7128, Longitude: -74.0060}
	assert.Equal(t, expectedValidInput, validInput, "TEST ok")

	// Invalid JSON
	invalidBody := bytes.NewBufferString(invalidJSON)
	_, err = ValidateInputData(invalidBody)
	assert.Error(t, err, "Error was expected")

	// Missing fields
	missingFieldsBody := bytes.NewBufferString(missingFieldsJSON)
	_, err = ValidateInputData(missingFieldsBody)
	assert.Error(t, err, "Expected error missing fields, test failed")
}
