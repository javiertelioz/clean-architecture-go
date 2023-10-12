package services

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

type JWTService interface {
	Generate(user *entity.User) (string, error)
	Verify(token string) (*entity.Token, error)
}
