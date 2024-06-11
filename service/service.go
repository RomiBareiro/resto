package service

import (
	"fmt"
	"resto_go/types"
)

// GetIDS returns merchant IDS that are available for the delivery
func GetIDS(in types.InputData, merchants []types.MerchantInfo) (types.Output, error) {
	if len(merchants) == 0 {
		return types.Output{}, fmt.Errorf("no available merchants")
	}
	var Ids []string
	for _, merchant := range merchants {
		if canDeliverHere(in.Latitude, in.Longitude, merchant.Latitude, merchant.Longitude, merchant.AvailabilityRadius) {
			Ids = append(Ids, merchant.ID)
		}
	}

	return types.Output{IDs: Ids}, nil
}
