package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	SigningKey     []byte
	Issuer         string
	ExpirationTime time.Duration
)

func NewToken(id string) (string, time.Time, error) {
	tokenID, err := uuid.NewV7()
	if err != nil {
		return "", time.Time{}, fmt.Errorf("jwt can't get new uuid v7: %w", err)
	}

	claims := jwt.RegisteredClaims{
		Issuer:    Issuer,
		Subject:   id,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpirationTime)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        tokenID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(SigningKey)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("jwt can't sign string: %w", err)
	}

	return ss, claims.ExpiresAt.Time, nil
}
