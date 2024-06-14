package service

import (
	"math"
	"time"
)

// canDeliverHere returns if it is deliverable (TRUE or FALSE)
func (s *service) canDeliverHere(deliveryLat, deliveryLon, customerLat, customerLon, availabilityRadius float64) bool {
	radius := 6371.0 //earth radius

	lat1 := toRadians(deliveryLat)
	lon1 := toRadians(deliveryLon)
	lat2 := toRadians(customerLat)
	lon2 := toRadians(customerLon)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	// Haversine
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := radius * c
	s.logger.Sugar().Infof("distance: %v\n", distance)
	return distance <= availabilityRadius
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
