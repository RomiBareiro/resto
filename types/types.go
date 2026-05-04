package types

import "time"

type MerchantInfo struct {
	ID                 string `csv:"id" db:"id"`
	Latitude           string `csv:"latitude" db:"latitude"`
	Longitude          string `csv:"longitude" db:"longitude"`
	AvailabilityRadius string `csv:"availability_radius" db:"availability_radius"`
	OpenHour           string `csv:"open_hour" db:"open_hour"`
	CloseHour          string `csv:"close_hour" db:"close_hour"`
	Rating             int    `csv:"rating" db:"rating"`
	OpenTime           time.Time
	CloseTime          time.Time
}

type InputData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Output struct {
	IDs []string `json:"ids,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Cause string `json:"cause,omitempty"`
}
