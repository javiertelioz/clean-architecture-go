package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"time"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() services.JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "pragmaticreviews.com",
	}
}

func getSecretKey() string {
	secret, _ := config.GetConfig[string]("Jwt.secret")

	return secret
}

func (jwtService jwtService) GenerateToken(username string, admin bool) string {
	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    jwtService.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtService.secretKey))

	if err != nil {
		panic(err)
	}

	return t
}

func (jwtService jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtService.secretKey), nil
	})
}
