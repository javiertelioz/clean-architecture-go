package user

import (
	"errors"
	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/user"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GetUsersUseCaseTestSuite struct {
	suite.Suite
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	users             []*entity.User
	result            []*entity.User
	err               error
}

func TestGetUsersUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetUsersUseCaseTestSuite))
}

func (suite *GetUsersUseCaseTestSuite) SetupTest() {
	suite.mockUserService = new(service.MockUserService)
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.users = []*entity.User{
		{
			ID:       1,
			LastName: "Doe",
			Name:     "John",
			Email:    "john@example.com",
			Phone:    "+123456789",
			Password: "password123",
		},
		{
			ID:       2,
			LastName: "Doe",
			Name:     "Jane",
			Email:    "jane@example.com",
			Phone:    "+123456790",
			Password: "password123",
		},
	}
}

func (suite *GetUsersUseCaseTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("GetUsers").Return(suite.users, nil)
}

func (suite *GetUsersUseCaseTestSuite) givenUserServiceReturnsError() {
	suite.mockUserService.On("GetUsers").Return(nil, errors.New("database error"))
}

func (suite *GetUsersUseCaseTestSuite) whenGetUsersUseCaseIsCalled() {
	suite.result, suite.err = user.GetUsersUseCase(suite.mockUserService, suite.mockLoggerService)
}

func (suite *GetUsersUseCaseTestSuite) thenExpectSuccess() {
	suite.NoError(suite.err)
	suite.NotNil(suite.result)
	suite.Equal(2, len(suite.result))
	suite.Equal(suite.users[0].Name, suite.result[0].Name)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUsersUseCaseTestSuite) thenExpectError() {
	suite.Error(suite.err)
	suite.Nil(suite.result)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUsersUseCaseTestSuite) TestGetUsersUseCaseWithSuccessResult() {
	// Given
	suite.givenUserServiceReturnsSuccess()

	// When
	suite.whenGetUsersUseCaseIsCalled()

	// Then
	suite.thenExpectSuccess()
}

func (suite *GetUsersUseCaseTestSuite) TestGetUsersUseCaseWithErrorResult() {
	// Given
	suite.givenUserServiceReturnsError()

	// When
	suite.whenGetUsersUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}
