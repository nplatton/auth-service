package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func init() {
	privKey := `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACBjn80LkCtCtAP8j1i4dJ4qJkaO7sAelMmFu1z6pBG9PQAAAJgbhV9FG4Vf
RQAAAAtzc2gtZWQyNTUxOQAAACBjn80LkCtCtAP8j1i4dJ4qJkaO7sAelMmFu1z6pBG9PQ
AAAEAKm/izgV7z7BWjejWw+BCCo2L8bJdHFRza5OD5dsj02GOfzQuQK0K0A/yPWLh0niom
Ro7uwB6UyYW7XPqkEb09AAAAEW5vYWhATnVnQm9vay5Ib21lAQIDBA==
-----END OPENSSH PRIVATE KEY-----`
	pubKey := `AAAAC3NzaC1lZDI1NTE5AAAAIGOfzQuQK0K0A/yPWLh0niomRo7uwB6UyYW7XPqkEb09`
}

func main() {
	fmt.Println("Hello, World!")
}

func GenerateToken(userID, role, secretKey string) (string, error) {
	claims := ComposeTokenClaims(userID, role)
	headers := ComposeTokenHeaders()
	token, err := SignToken(secretKey, headers, claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ComposeTokenHeaders() jwt.Header {
	headers := jwt.Header{
		"alg": "HS256",
		"typ": "JWT",
		"kid": "auth-service-2026",
	}
	
	return headers
}

func ComposeTokenClaims(userID, role string) jwt.MapClaims {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"iat":	 jwt.NewNumericDate(jwt.TimeFunc()),
		"exp":	 jwt.NewNumericDate(jwt.TimeFunc().Add(24 * 60 * 60 * 1e9)), // Token expires in 24 hours
		"nbf":	 jwt.NewNumericDate(jwt.TimeFunc()), // Not before now
		"aud":	 "finance-app",
		"iss":	 "auth-service",
		"sub":	 userID,
		"jti":	 GenerateJTI(),
	}
	
	return claims
}

func GenerateJTI() string {
	// Generate a unique token ID (JTI) using a random string or UUID
	// For simplicity, we'll use a timestamp-based approach here
	return fmt.Sprintf("%d", jwt.TimeFunc().UnixNano())
}

func SignToken(secretKey string, headers jwt.Header, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header = headers
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}