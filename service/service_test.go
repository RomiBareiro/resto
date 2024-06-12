package service

import (
	"reflect"
	"resto_go/types"
	"testing"
	"time"
)

func TestGetIDs(t *testing.T) {
	// Testing data
	in := types.InputData{Latitude: 40.7128, Longitude: -74.0060}
	merchants := []types.MerchantInfo{
		{ID: "1", Latitude: 40.71, Longitude: -74.01, AvailabilityRadius: 5.0, OpenHour: time.Now(), CloseHour: time.Now().Add(1 * time.Hour)},
		{ID: "2", Latitude: 60.72, Longitude: -84.02, AvailabilityRadius: 3.0, OpenHour: time.Now(), CloseHour: time.Now().Add(1 * time.Hour)},
		{ID: "3", Latitude: 40.71, Longitude: -74.02, AvailabilityRadius: 5.0, OpenHour: time.Now(), CloseHour: time.Now().Add(1 * time.Hour)},
	}

	output, err := GetIDS(in, merchants)

	if err != nil {
		t.Errorf("Error in GetIDs: %v", err)
	}

	expectedOutput := types.Output{IDs: []string{"1", "3"}}
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Unexpected output. Expected: %v, got: %v", expectedOutput, output)
	}
}
