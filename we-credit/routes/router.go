package routes

import (
	"otp-auth/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	controllers.InitController(db)

	// Register User
	r.POST("/register", controllers.RegisterUser)

	// Login User
	r.POST("/login", controllers.LoginUser)

	// Generate OTP
	r.POST("/generate-otp", controllers.GenerateOTP)

	// Verify OTP
	r.POST("/verify-otp", controllers.VerifyOTP)
}
