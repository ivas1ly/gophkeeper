package jwt

import (
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	SigningKey = []byte("2818395a611a724b560d932274f43d2a3ebc9042929fabe5c2f91db185c25adb")
	Issuer = "gophkeeper"
	ExpirationTime = 12 * time.Hour

	id, err := uuid.NewV7()
	assert.NoError(t, err)

	t.Run("check token", func(t *testing.T) {
		signedToken, expiredAt, err := NewToken(id.String())
		assert.NoError(t, err)
		assert.Equal(t, len(strings.Split(signedToken, ".")), 3)

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(signedToken, claims, func(_ *jwt.Token) (interface{}, error) {
			return SigningKey, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, token.Valid, true)
		assert.Equal(t, claims["sub"], id.String())
		assert.Equal(t, time.Now().Before(expiredAt), true)
	})
}
