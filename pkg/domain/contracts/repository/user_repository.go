package repository

import "github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id int64) error
	FindByEmail(email string) (*entity.User, error)
}
