package types

import (
	"time"
)

type MerchantInfo struct {
	ID                 string    `csv:"id"`
	Latitude           float64   `csv:"latitude"`
	Longitude          float64   `csv:"longitude"`
	AvailabilityRadius float64   `csv:"availability_radius"`
	OpenHour           time.Time `csv:"open_hour" format:"2006-01-02T15:04:05Z"`
	CloseHour          time.Time `csv:"close_hour" format:"2006-01-02T15:04:05Z"`
	Rating             int       `csv:"rating"`
}

type InputData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Output struct {
	IDs []string `json:"ids"`
}
