package service

import (
	"reflect"
	"resto_go/types"
	"testing"
)

func TestGetIDS(t *testing.T) {
	// Testing data
	in := types.InputData{Latitude: 40.7128, Longitude: -74.0060}
	merchants := []types.MerchantInfo{
		{ID: "1", Latitude: 40.71, Longitude: -74.01, AvailabilityRadius: 5.0},
		{ID: "2", Latitude: 60.72, Longitude: -84.02, AvailabilityRadius: 3.0}, // out of availability radius
		{ID: "3", Latitude: 40.72, Longitude: -74.02, AvailabilityRadius: 3.0},
	}

	output, err := GetIDS(in, merchants)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := types.Output{IDs: []string{"1", "3"}}
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Unexpected output. Expected: %v, got: %v", expectedOutput, output)
	}
}
