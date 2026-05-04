package service

import (
	"fmt"
	t "resto_go/types"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// Service define la interfaz para nuestro servicio
type Service interface {
	GetIDS(in t.InputData, merchants []t.MerchantInfo) (*t.Output, error)
	IsMerchantOpen(openTime, closeTime time.Time) bool
	Pool() *pgxpool.Pool
}

type service struct {
	logger *zap.Logger
	DB     *pgxpool.Pool
}

func NewService(logger *zap.Logger, conn *pgxpool.Pool) Service {
	return &service{logger: logger, DB: conn}
}
func (s *service) Pool() *pgxpool.Pool {
	return s.DB
}

// GetIDS returns merchant IDS that are available for the delivery
func (s *service) GetIDS(in t.InputData, merchants []t.MerchantInfo) (*t.Output, error) {
	if len(merchants) == 0 {
		return nil, fmt.Errorf("no available merchants")
	}
	var Ids []string
	for _, merchant := range merchants {
		if s.canDeliverHere(in, merchant) {
			if s.IsMerchantOpen(merchant.OpenTime, merchant.CloseTime) {
				Ids = append(Ids, merchant.ID)
			}
		}
	}
	if len(Ids) == 0 {
		return nil, fmt.Errorf("no available merchants")
	}
	return &t.Output{IDs: Ids}, nil
}
