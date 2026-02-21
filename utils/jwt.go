package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("SECRET_KEY_CHANGE_THIS") // you can change this to any random string

type JWTClaim struct {
    UserID string `json:"user_id"`
    Email  string `json:"email"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

// GenerateJWT creates a new token for a user
func GenerateJWT(userID, email, role string) (string, error) {
    claims := &JWTClaim{
        UserID: userID,
        Email:  email,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 1 day expiry
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateToken parses and validates a JWT token
func ValidateToken(tokenString string) (*JWTClaim, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
        return claims, nil
    } else {
        return nil, err
    }
}