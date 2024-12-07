package adapter

import "github.com/golang-jwt/jwt/v5"

type JwtEchoAdapter struct {
	
}

func NewJwtAdapter() *JwtEchoAdapter {
	return &JwtEchoAdapter{}
}

func (ad *JwtEchoAdapter) Generate() (string) {
	token := jwt.New(jwt.SigningMethodHS256)
	return token.Raw
}