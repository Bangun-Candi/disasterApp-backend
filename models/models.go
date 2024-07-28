package models

import "time"

type Users struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// type RescueRequest struct {
// 	UserEmail          string        `json:"userEmail"`
// 	UserName           string        `json:"userName"`
// 	UserPhoneNumber    string        `json:"userPhoneNumber"`
// 	LongitudeLocation  string        `json:"longitudeLocation"`
// 	LatitudeLocation   string        `json:"latitudeLocation"`
// 	RescuersCode       []RescuerCode `json:"rescuersCode"`
// 	DisasterTypeCode   string        `json:"disasterTypeCode"`
// 	StatusDisaster     string        `json:"statusDisaster"`
// 	StatusDisasterCode string        `json:"statusDisasterCode"`
// 	Notes              string        `json:"notes"`
// }

// type Rescuer struct {
// 	ID           int    `json:"id"`
// 	CategoryName string `json:"category_name"`
// 	CategoryCode string `json:"category_code"`
// }

// type RescueRequest struct {
// 	UserEmail         string       `json:"userEmail"`
// 	UserName          string       `json:"userName"`
// 	UserPhoneNumber   string       `json:"userPhoneNumber"`
// 	LongitudeLocation string       `json:"longitudeLocation"`
// 	LatitudeLocation  string       `json:"latitudeLocation"`
// 	RescuersCode      []RescuerCode `json:"rescuersCode"`
// 	DisasterTypeCode  string       `json:"disasterTypeCode"`
// 	StatusDisaster    string       `json:"statusDisaster"`
// 	StatusDisasterCode string      `json:"statusDisasterCode"`
// 	Notes             string       `json:"notes"`
// }

// type Disaster struct {
// 	ID           int       `json:"id"`
// 	DisasterName string    `json:"disaster_name"`
// 	DisasterCode string    `json:"disaster_code"`
// 	StatusCode   string    `json:"status_code"`
// 	Latitude     float64   `json:"latitude"`
// 	Longitude    float64   `json:"longitude"`
// 	Notes        string    `json:"notes"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// }

// type RealTimeLocation struct {
// 	ID           int       `json:"id"`
// 	UserID       int       `json:"user_id"`
// 	Latitude     float64   `json:"latitude"`
// 	Longitude    float64   `json:"longitude"`
// 	LocationName string    `json:"location_name"`
// 	DisasterCode string    `json:"disaster_code"`
// 	DisasterName string    `json:"disaster_name"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// }
