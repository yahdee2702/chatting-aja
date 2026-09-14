package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const ISSUER_NAME = "CHATTING-AJA-API"
const EXPIRATON_DURATION = 24 * 3 * time.Hour // 3 Days

type JwtHandler struct {
	secret []byte
}

func NewJwtHandler(secret []byte) *JwtHandler {
	return &JwtHandler{
		secret: secret,
	}
}

func (j *JwtHandler) Generate(userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		Issuer:    "API",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(EXPIRATON_DURATION)),
	}

	token := jwt.NewWithClaims(&jwt.SigningMethodHMAC{}, claims)

	return token.SignedString(j.secret)
}

func (j *JwtHandler) Verify(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return j.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, nil
}
