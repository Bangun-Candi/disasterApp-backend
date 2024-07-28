package models

import "time"

type OTPCode struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	OTPCode     string    `json:"otp_code"`
	Type        string    `json:"type"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
}
