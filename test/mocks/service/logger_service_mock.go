package service

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockLoggerService struct {
	mock.Mock
}

func (ls *MockLoggerService) Trace(msg string) {
	fmt.Println(msg)
}

func (ls *MockLoggerService) Info(msg string) {
	fmt.Println(msg)
}

func (ls *MockLoggerService) Debug(msg string) {
	fmt.Println(msg)
}

func (ls *MockLoggerService) Warn(msg string) {
	fmt.Println(msg)
}

func (ls *MockLoggerService) Error(msg string) {
	fmt.Println(msg)
}
