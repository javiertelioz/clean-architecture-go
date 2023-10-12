package services

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

type UserService interface {
	GetUser(id string) (*entity.User, error)
	GetUsers() ([]*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id string) error
	GetUserByEmail(email string) (*entity.User, error)
}
