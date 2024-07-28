package models

import "time"

// type User struct {
// 	ID             int    `json:"id"`
// 	Email          string `json:"email"`
// 	PhoneNumber    string `json:"phone_number"`
// 	Username       string `json:"username"`
// 	NIK            string `json:"nik"`
// 	Name           string `json:"name"`
// 	BirthdayDate   string `json:"birthday_date"`
// 	BirthdayPlace  string `json:"birthday_place"`
// 	Gender         string `json:"gender"`
// 	Address        string `json:"address"`
// 	Religion       string `json:"religion"`
// 	MarriageStatus string `json:"marriage_status"`
// 	Occupation     string `json:"occupation"`
// 	Citizenship    string `json:"citizenship"`
// 	ValidityPeriod string `json:"validity_period"`
// 	Password       string `json:"password"`
// 	PIN            string `json:"pin"`
// }

type User struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RescueRequest struct {
	UserEmail          string        `json:"userEmail"`
	UserName           string        `json:"userName"`
	UserPhoneNumber    string        `json:"userPhoneNumber"`
	LongitudeLocation  string        `json:"longitudeLocation"`
	LatitudeLocation   string        `json:"latitudeLocation"`
	RescuersCode       []RescuerCode `json:"rescuersCode"`
	DisasterTypeCode   string        `json:"disasterTypeCode"`
	StatusDisaster     string        `json:"statusDisaster"`
	StatusDisasterCode string        `json:"statusDisasterCode"`
	Notes              string        `json:"notes"`
}

type RescuerCode struct {
	Code string `json:"code"`
}

type Rescuer struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CategoryCode string `json:"category_code"`
}

type Disaster struct {
	ID           int       `json:"id"`
	DisasterName string    `json:"disaster_name"`
	DisasterCode string    `json:"disaster_code"`
	StatusCode   string    `json:"status_code"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RealTimeLocation struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	LocationName string    `json:"location_name"`
	DisasterCode string    `json:"disaster_code"`
	DisasterName string    `json:"disaster_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
