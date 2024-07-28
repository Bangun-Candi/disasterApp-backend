package models

import "time"

type KYCData struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	NIK            string    `json:"nik"`
	Name           string    `json:"name"`
	BirthdayDate   time.Time `json:"birthday_date"`
	BirthdayPlace  string    `json:"birthday_place"`
	Gender         string    `json:"gender"`
	Address        string    `json:"address"`
	Religion       string    `json:"religion"`
	MarriageStatus string    `json:"marriage_status"`
	Occupation     string    `json:"occupation"`
	Citizenship    string    `json:"citizenship"`
	ValidityPeriod time.Time `json:"validity_period"`
	CreatedAt      time.Time `json:"created_at"`
}
