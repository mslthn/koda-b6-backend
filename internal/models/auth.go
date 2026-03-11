package models

import "time"

type ForgotPassword struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	OTPCode   string    `json:"otp_code"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
