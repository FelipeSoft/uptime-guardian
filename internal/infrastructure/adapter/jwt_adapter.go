package adapter

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAdapter struct {
	Secret string
}

func NewJwtAdapter() *JwtAdapter {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not set in environment variables")
	}
	return &JwtAdapter{Secret: secret}
}

func (ad *JwtAdapter) Generate(metadata string) (string, error) {
	claims := jwt.MapClaims{
		"metadata": metadata,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
		"iat":      time.Now().Unix(),                    
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(ad.Secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (ad *JwtAdapter) Expired(tokenString string) bool {
	return true
}

func (ad *JwtAdapter) Read(tokenString string) (*string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("not allowed token method")
		}
		return []byte(ad.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return &token.Raw, nil
}
