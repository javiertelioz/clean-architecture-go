package user

import (
	"errors"
	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/user"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/test/application/user_cases/user/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CreateUserUseCaseTestSuite struct {
	suite.Suite
	mockUserService   *mocks.MockUserService
	mockLoggerService *mocks.MockLoggerService
	user              *entity.User
	result            *entity.User
	err               error
}

func TestCreateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUseCaseTestSuite))
}

func (suite *CreateUserUseCaseTestSuite) SetupTest() {
	suite.mockUserService = new(mocks.MockUserService)
	suite.mockLoggerService = new(mocks.MockLoggerService)
	suite.user = &entity.User{
		ID:       1,
		LastName: "Doe",
		Name:     "John",
		Email:    "john@example.com",
		Phone:    "+123456789",
		Password: "password123",
	}
}

func (suite *CreateUserUseCaseTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("CreateUser", suite.user).Return(suite.user, nil)
}

func (suite *CreateUserUseCaseTestSuite) givenUserServiceReturnsError() {
	suite.mockUserService.On("CreateUser", suite.user).Return(nil, errors.New("database error"))
}

func (suite *CreateUserUseCaseTestSuite) whenCreateUserUseCaseIsCalled() {
	suite.result, suite.err = user.CreateUserUseCase(suite.user, suite.mockUserService, suite.mockLoggerService)
}

func (suite *CreateUserUseCaseTestSuite) thenExpectSuccess() {
	suite.NoError(suite.err)
	suite.NotNil(suite.result)
	suite.Equal(suite.user.Name, suite.result.Name)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *CreateUserUseCaseTestSuite) thenExpectError() {
	suite.Error(suite.err)
	suite.Nil(suite.result)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUserUseCaseWithSuccessResult() {
	// Given
	suite.givenUserServiceReturnsSuccess()

	// When
	suite.whenCreateUserUseCaseIsCalled()

	// Then
	suite.thenExpectSuccess()
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUserUseCaseWithErrorResult() {
	// Given
	suite.givenUserServiceReturnsError()

	// When
	suite.whenCreateUserUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}
