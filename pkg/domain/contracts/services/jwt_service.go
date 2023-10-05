package services

import "github.com/dgrijalva/jwt-go"

type JWTService interface {
	GenerateToken(username string, admin bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
