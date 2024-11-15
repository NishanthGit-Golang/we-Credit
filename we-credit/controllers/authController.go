package controllers

import (
	"log"
	"net/http"
	"otp-auth/models"
	"otp-auth/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

// Initialize the db connection for controllers
func InitController(database *gorm.DB) {
	db = database
}

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var request struct {
		MobileNumber      string `json:"mobile_number"`
		DeviceFingerprint string `json:"device_fingerprint"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Create user in database
	user := models.User{
		MobileNumber:      request.MobileNumber,
		DeviceFingerprint: request.DeviceFingerprint,
	}

	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var request struct {
		MobileNumber      string `json:"mobile_number"`
		DeviceFingerprint string `json:"device_fingerprint"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if user exists
	var user models.User
	if err := db.Where("mobile_number = ? AND device_fingerprint = ?", request.MobileNumber, request.DeviceFingerprint).First(&user).Error; err != nil {
		log.Printf("User not found: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}

// GenerateOTP generates a new OTP for a user
func GenerateOTP(c *gin.Context) {
	var request struct {
		MobileNumber string `json:"mobile_number"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Generate OTP and save it in DB
	otp, err := services.CreateOTP(db, request.MobileNumber)
	if err != nil {
		log.Printf("Error creating OTP: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"otp": otp.Code})
}

// VerifyOTP verifies the OTP entered by the user
func VerifyOTP(c *gin.Context) {
	var request struct {
		MobileNumber string `json:"mobile_number"`
		OTP          string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Retrieve the latest OTP
	otp, err := services.GetLatestOTP(db, request.MobileNumber)
	if err != nil {
		log.Printf("Error retrieving OTP: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "OTP not found"})
		return
	}

	// Check if OTP is expired
	if services.IsOTPExpired(otp) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "OTP expired"})
		return
	}

	// Verify OTP
	if otp.Code != request.OTP {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
