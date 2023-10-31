package service

import (
	"github.com/stretchr/testify/mock"
)

type MockCryptoService struct {
	mock.Mock
}

func (m *MockCryptoService) Hash(password string) (string, error) {
	args := m.Called(password)

	if args.Get(0) == nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}

func (m *MockCryptoService) Verify(password, hashedPassword string) error {
	args := m.Called(password, hashedPassword)
	return args.Error(0)
}
