package token

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/hipeday/upay/internal/constants"
	"time"
)

// GenerateToken generates a random token of the specified length.
func GenerateToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GenerateAccessToken generates an access token with expiration.
func GenerateAccessToken() (string, time.Time, error) {
	token, err := GenerateToken(32) // 32 bytes = 64 hex characters
	if err != nil {
		return "", time.Time{}, err
	}
	expiresAt := time.Now().Add(time.Hour * constants.TokenValidityPeriod) // Access token valid for 1 hour
	return token, expiresAt, nil
}

// GenerateRefreshToken generates a refresh token with expiration.
func GenerateRefreshToken() (string, error) {
	token, err := GenerateToken(64) // 64 bytes = 128 hex characters
	if err != nil {
		return "", err
	}
	// refresh 不设置过期时间
	//expiresAt := time.Now().Add(time.Hour * 24 * 7) // Refresh token valid for 7 days
	return token, nil
}
