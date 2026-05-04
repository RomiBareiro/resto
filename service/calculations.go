package service

import (
	"math"
	t "resto_go/types"
	"strconv"
	"time"
)

// canDeliverHere returns if it is deliverable (TRUE or FALSE)
func (s *service) canDeliverHere(in t.InputData, merchantInfo t.MerchantInfo) bool {
	radius := 6371.0 //earth radius

	// Convert string coordinates to float64
	lat2, err := strconv.ParseFloat(merchantInfo.Latitude, 64)
	if err != nil {
		s.logger.Sugar().Errorf("failed to parse latitude: %v", err)
		return false
	}
	lon2, err := strconv.ParseFloat(merchantInfo.Longitude, 64)
	if err != nil {
		s.logger.Sugar().Errorf("failed to parse longitude: %v", err)
		return false
	}
	availRadius, err := strconv.ParseFloat(merchantInfo.AvailabilityRadius, 64)
	if err != nil {
		s.logger.Sugar().Errorf("failed to parse availability radius: %v", err)
		return false
	}

	lat1 := toRadians(in.Latitude)
	lon1 := toRadians(in.Longitude)
	lat2Rad := toRadians(lat2)
	lon2Rad := toRadians(lon2)

	dLat := lat2Rad - lat1
	dLon := lon2Rad - lon1

	// Haversine
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2Rad)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := radius * c
	s.logger.Sugar().Infof("distance: %v\n availability radius: %v", distance, availRadius)
	return distance <= availRadius
}

// toRadians convert degrees to radians
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (s *service) IsMerchantOpen(openTime, closeTime time.Time) bool {
	currentTime := time.Now()

	if currentTime.After(openTime) && currentTime.Before(closeTime) {
		// faltan mas de 10 min para cerrar y esta abierto
		timeUntilClosing := closeTime.Sub(currentTime)
		s.logger.Sugar().Infof("merchant is open, remaining time: %v", timeUntilClosing)
		return timeUntilClosing > 10*time.Minute
	}
	return false
}
