package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing JWT
var jwtSecret = []byte("CDy3Lo7iOhIbeRP9RDDN5I6cqp1ZLCdb2EFlR2/D3Gg=")

// Claims structure
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
