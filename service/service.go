package service

import (
	"fmt"
	"resto_go/types"
	"time"

	"go.uber.org/zap"
)

// Service define la interfaz para nuestro servicio
type Service interface {
	GetIDS(in types.InputData, merchants []types.MerchantInfo) (types.Output, error)
	IsMerchantOpen(openTime, closeTime time.Time) bool
}

type service struct {
	logger *zap.Logger
}

func NewService(logger *zap.Logger) Service {
	return &service{logger: logger}
}

// GetIDS returns merchant IDS that are available for the delivery
func (s *service) GetIDS(in types.InputData, merchants []types.MerchantInfo) (types.Output, error) {
	if len(merchants) == 0 {
		return types.Output{}, fmt.Errorf("no available merchants")
	}
	var Ids []string
	for _, merchant := range merchants {
		if s.canDeliverHere(in.Latitude, in.Longitude, merchant.Latitude, merchant.Longitude, merchant.AvailabilityRadius) {
			if s.IsMerchantOpen(merchant.OpenHour, merchant.CloseHour) {
				Ids = append(Ids, merchant.ID)
			}
		}
	}

	return types.Output{IDs: Ids}, nil
}
