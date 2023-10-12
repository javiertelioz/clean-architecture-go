package services

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
	salt int
}

func NewBcryptService(salt int) services.CryptoService {
	return &BcryptService{
		salt: salt,
	}
}

func (bs BcryptService) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bs.salt)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (bs BcryptService) Verify(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
