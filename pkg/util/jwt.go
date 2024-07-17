package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	AccountId int64 `json:"account_id"`
	jwt.StandardClaims
}

// ParseToken 解析JWT
func ParseToken(tokenString string, secret string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	return token, claims, err
}

// ValidateToken validates the token.
func ValidateToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// GenerateToken generates a token.
func GenerateToken(accountId int64, secret string, expiresAt *time.Duration) (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	standardClaims := jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Id:       uid.String(),
	}

	if expiresAt != nil {
		expirationTime := time.Now().Add(*expiresAt)
		standardClaims.ExpiresAt = expirationTime.Unix()
	}

	claims := &Claims{
		AccountId:      accountId,
		StandardClaims: standardClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}
