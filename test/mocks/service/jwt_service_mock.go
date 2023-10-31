package service

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockJwtService struct {
	mock.Mock
}

func (js *MockJwtService) Generate(user *entity.User) (string, error) {
	args := js.Called(user)

	if args.Get(0) == nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}

func (js *MockJwtService) Verify(token string) (*entity.Token, error) {
	args := js.Called(token)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.Token), args.Error(1)
}
