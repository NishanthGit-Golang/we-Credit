package models

type OTP struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	MobileNumber string `gorm:"index" json:"mobile_number"`
	Code         string `json:"code"`
	ExpiresAt    int64  `json:"expires_at"`
}
