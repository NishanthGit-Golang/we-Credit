package main

import (
	"log"
	"otp-auth/models"
	"otp-auth/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// Setup MySQL connection
	dsn := "root:Nishanth123#@tcp(127.0.0.1:3306)/choiceDB?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Apply AutoMigration for User and OTP tables
	if err := db.AutoMigrate(&models.User{}, &models.OTP{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}

func main() {
	r := gin.Default()

	// Initialize routes
	routes.SetupRoutes(r, db)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
