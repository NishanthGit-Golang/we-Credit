package models

type User struct {
	ID                uint   `gorm:"primary_key" json:"id"`
	MobileNumber      string `gorm:"unique;not null" json:"mobile_number"`
	DeviceFingerprint string `json:"device_fingerprint"`
}
