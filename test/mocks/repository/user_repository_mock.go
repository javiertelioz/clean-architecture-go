package repository

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user *entity.User) (*entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(id string) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) FindAll() ([]*entity.User, error) {
	args := m.Called()
	return args.Get(0).([]*entity.User), args.Error(1)
}

func (m *MockUserRepository) Update(user *entity.User) (*entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) Delete(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}
