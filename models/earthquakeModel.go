package models

import "time"

type Earthquake struct {
	ID             int       `json:"id"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Magnitude      float64   `json:"magnitude"`
	Depth          float64   `json:"depth"`
	Location       string    `json:"location"`
	OccurrenceTime time.Time `json:"occurrence_time"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
