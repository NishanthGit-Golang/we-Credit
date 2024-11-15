package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateFingerprint(userAgent string, ipAddress string) string {
	hash := sha256.New()
	hash.Write([]byte(userAgent + ipAddress))
	return hex.EncodeToString(hash.Sum(nil))
}
