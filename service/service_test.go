package service

import (
	"reflect"
	"resto_go/types"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestGetIDs(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	svc := NewService(logger)

	// Testing data
	in := types.InputData{Latitude: 40.7128, Longitude: -74.0060}
	merchants := []types.MerchantInfo{
		{ID: "1", Latitude: 40.71, Longitude: -74.01, AvailabilityRadius: 5.0, OpenHour: time.Now(), CloseHour: time.Now().Add(1 * time.Hour)},
		{ID: "2", Latitude: 60.72, Longitude: -84.02, AvailabilityRadius: 3.0, OpenHour: time.Now(), CloseHour: time.Now().Add(1 * time.Hour)},
		{ID: "3", Latitude: 40.71, Longitude: -74.02, AvailabilityRadius: 5.0, OpenHour: time.Now(), CloseHour: time.Now().Add(1 * time.Hour)},
	}

	output, err := svc.GetIDS(in, merchants)

	if err != nil {
		t.Errorf("Error in GetIDs: %v", err)
	}

	expectedOutput := types.Output{IDs: []string{"1", "3"}}
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Unexpected output. Expected: %v, got: %v", expectedOutput, output)
	}
}

func TestIsMerchantOpen(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	svc := NewService(logger)
	tests := []struct {
		name      string
		openTime  time.Time
		closeTime time.Time
		expected  bool
	}{
		{
			name:      "Open but closing soon",
			openTime:  time.Now().Add(-2 * time.Hour),
			closeTime: time.Now().Add(10 * time.Minute),
			expected:  false,
		},
		{
			name:      "Open and not closing soon",
			openTime:  time.Now().Add(-2 * time.Hour),
			closeTime: time.Now().Add(2 * time.Hour),
			expected:  true,
		},
		{
			name:      "Not open yet",
			openTime:  time.Now().Add(1 * time.Hour),
			closeTime: time.Now().Add(3 * time.Hour),
			expected:  false,
		},
		{
			name:      "Already closed",
			openTime:  time.Now().Add(-3 * time.Hour),
			closeTime: time.Now().Add(-2 * time.Hour),
			expected:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := svc.IsMerchantOpen(test.openTime, test.closeTime)
			if result != test.expected {
				t.Errorf("Expected %v but got %v for test case: %s", test.expected, result, test.name)
			}
		})
	}
}
