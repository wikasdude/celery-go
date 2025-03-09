package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key (should be the same as in auth.go)
var jwtSecret = []byte("CDy3Lo7iOhIbeRP9RDDN5I6cqp1ZLCdb2EFlR2/D3Gg=")

// Claims structure
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// VerifyJWT extracts and verifies the token from the request
func VerifyJWT(r *http.Request) (*Claims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("no token provided")
	}

	// Extract the token from "Bearer <token>"
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
