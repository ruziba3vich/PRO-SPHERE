package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define the Claims struct as per your question
type Claims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
	ProID  int `json:"pro_id"`

	// User *models.User
}

// Secret key for signing the token (keep this secret in production)

// Function to generate JWT token
func GenerateJWT(userID, proID int, jwtSecret []byte) (string, error) {
	// Set the claims

	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(), // Token expires in 24 hours
			Issuer:    "pro-sphere",                          // Set your issuer here
		},
		UserID: userID,
		ProID:  proID,
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return signedToken, nil
}

// Function to extract claims from the JWT token
func ExtractClaims(tokenString string, jwtSecret []byte) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the token method and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
