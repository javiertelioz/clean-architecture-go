package mocks

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUser(id string) (*entity.User, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserService) GetUsers() ([]*entity.User, error) {
	args := m.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*entity.User), args.Error(1)
}

func (m *MockUserService) CreateUser(user *entity.User) (*entity.User, error) {
	args := m.Called(user)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(user *entity.User) (*entity.User, error) {
	args := m.Called(user)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
