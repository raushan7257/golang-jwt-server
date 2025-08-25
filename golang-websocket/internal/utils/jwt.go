package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// Generate random secret
func GenerateSecret() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
