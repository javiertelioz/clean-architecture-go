package services

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/repository"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
)

type UserService struct {
	repository repository.UserRepository
	logger     services.LoggerService
}

func NewUserService(
	repository repository.UserRepository,
	logger services.LoggerService,
) services.UserService {
	return &UserService{
		repository: repository,
		logger:     logger,
	}
}

func (s *UserService) GetUsers() ([]*entity.User, error) {
	return s.repository.FindAll()
}

func (s *UserService) GetUser(id string) (*entity.User, error) {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return nil, exceptions.UserNotFound()
	}

	return user, nil
}

func (s *UserService) CreateUser(user *entity.User) (*entity.User, error) {
	return s.repository.Create(user)
}

func (s *UserService) UpdateUser(user *entity.User) (*entity.User, error) {
	updateUser, err := s.repository.Update(user)

	if err != nil {
		return nil, exceptions.UserNotFound()
	}

	return updateUser, nil
}

func (s *UserService) DeleteUser(id string) error {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return exceptions.UserNotFound()
	}

	return s.repository.Delete(int64(user.ID))
}

func (s *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return nil, exceptions.UserNotFound()
	}

	return user, nil
}
