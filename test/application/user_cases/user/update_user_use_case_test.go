package user

import (
	"errors"
	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/user"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UpdateUserUseCaseTestSuite struct {
	suite.Suite
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	user              *entity.User
	result            *entity.User
	err               error
}

func TestUpdateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateUserUseCaseTestSuite))
}

func (suite *UpdateUserUseCaseTestSuite) SetupTest() {
	suite.mockUserService = new(service.MockUserService)
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.user = &entity.User{
		ID:       1,
		LastName: "Doe",
		Name:     "John",
		Email:    "john@example.com",
		Phone:    "+123456789",
		Password: "password123",
	}
}

func (suite *UpdateUserUseCaseTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("UpdateUser", suite.user).Return(suite.user, nil)
}

func (suite *UpdateUserUseCaseTestSuite) givenUserServiceReturnsError() {
	suite.mockUserService.On("UpdateUser", suite.user).Return(nil, errors.New("database error"))
}

func (suite *UpdateUserUseCaseTestSuite) whenUpdateUserUseCaseIsCalled() {
	suite.result, suite.err = user.UpdateUserUseCase(suite.user, suite.mockUserService, suite.mockLoggerService)
}

func (suite *UpdateUserUseCaseTestSuite) thenExpectSuccess() {
	suite.NoError(suite.err)
	suite.NotNil(suite.result)
	suite.Equal(suite.user.Name, suite.result.Name)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *UpdateUserUseCaseTestSuite) thenExpectError() {
	suite.Error(suite.err)
	suite.Nil(suite.result)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *UpdateUserUseCaseTestSuite) TestCreateUserUseCaseWithSuccessResult() {
	// Given
	suite.givenUserServiceReturnsSuccess()

	// When
	suite.whenUpdateUserUseCaseIsCalled()

	// Then
	suite.thenExpectSuccess()
}

func (suite *UpdateUserUseCaseTestSuite) TestCreateUserUseCaseWithErrorResult() {
	// Given
	suite.givenUserServiceReturnsError()

	// When
	suite.whenUpdateUserUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}
