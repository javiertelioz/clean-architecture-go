package user

import (
	"testing"

	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/user"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	user2 "github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
	"github.com/stretchr/testify/suite"
)

type GetUserByIdUseCaseTestSuite struct {
	suite.Suite
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	userID            string
	user              *entity.User
	result            *entity.User
	err               error
}

func TestGetUserByIdUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserByIdUseCaseTestSuite))
}

func (suite *GetUserByIdUseCaseTestSuite) SetupTest() {
	suite.mockUserService = new(service.MockUserService)
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.userID = "1"
	suite.user = &entity.User{
		ID:       1,
		LastName: "Doe",
		Name:     "John",
		Email:    "john@example.com",
		Phone:    "+123456789",
		Password: "password123",
	}
}

func (suite *GetUserByIdUseCaseTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("GetUser", suite.userID).Return(suite.user, nil)
}

func (suite *GetUserByIdUseCaseTestSuite) givenUserServiceReturnsError() {
	suite.mockUserService.On("GetUser", suite.userID).Return(nil, user2.UserNotFound())
}

func (suite *GetUserByIdUseCaseTestSuite) whenGetUserByIdUseCaseIsCalled() {
	suite.result, suite.err = user.GetUserByIdUseCase(suite.userID, suite.mockUserService, suite.mockLoggerService)
}

func (suite *GetUserByIdUseCaseTestSuite) thenExpectSuccess() {
	suite.NoError(suite.err)
	suite.NotNil(suite.result)
	suite.Equal(suite.user.Name, suite.result.Name)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUserByIdUseCaseTestSuite) thenExpectError() {
	suite.Error(suite.err)
	suite.Nil(suite.result)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUserByIdUseCaseTestSuite) TestGetUserByIdUseCaseWithSuccessResult() {
	// Given
	suite.givenUserServiceReturnsSuccess()

	// When
	suite.whenGetUserByIdUseCaseIsCalled()

	// Then
	suite.thenExpectSuccess()
}

func (suite *GetUserByIdUseCaseTestSuite) TestGetUserByIdUseCaseWithErrorResult() {
	// Given
	suite.givenUserServiceReturnsError()

	// When
	suite.whenGetUserByIdUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}
