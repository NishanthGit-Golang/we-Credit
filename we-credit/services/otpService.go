package services

import (
	"otp-auth/models"
	"time"

	"gorm.io/gorm"
)

// CreateOTP generates and saves a new OTP for a user
func CreateOTP(db *gorm.DB, mobileNumber string) (*models.OTP, error) {
	// Generate OTP (hardcoded for simplicity, you can implement a random generator here)
	otpCode := "123456"

	// Calculate expiration time (5 minutes)
	expiresAt := time.Now().Add(5 * time.Minute).Unix()

	// Create new OTP record
	otp := models.OTP{
		MobileNumber: mobileNumber,
		Code:         otpCode,
		ExpiresAt:    expiresAt,
	}

	// Save OTP to database
	if err := db.Create(&otp).Error; err != nil {
		return nil, err
	}

	return &otp, nil
}

// IsOTPExpired checks if the OTP is expired
func IsOTPExpired(otp *models.OTP) bool {
	return time.Now().Unix() > otp.ExpiresAt
}

// GetLatestOTP retrieves the latest OTP for a given mobile number
func GetLatestOTP(db *gorm.DB, mobileNumber string) (*models.OTP, error) {
	var otp models.OTP
	if err := db.Where("mobile_number = ?", mobileNumber).Order("created_at desc").First(&otp).Error; err != nil {
		return nil, err
	}
	return &otp, nil
}
