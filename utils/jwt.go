package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"errors"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type JWTClaim struct {
    UserID uint
    jwt.StandardClaims
}

func GenerateJWT(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
    token, err := jwt.ParseWithClaims(
        signedToken,
        &JWTClaim{},
        func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        },
    )
    if err != nil {
        return nil, err
    }
    claims, ok := token.Claims.(*JWTClaim)
    if !ok {
        return nil, errors.New("couldn't parse claims")
    }
    if claims.ExpiresAt < time.Now().Unix() {
        return nil, errors.New("token expired")
    }
    return claims, nil
}