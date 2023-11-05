package user

import (
	"testing"

	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/user"
	user2 "github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
	"github.com/stretchr/testify/suite"
)

type DeleteUserUseCaseTestSuite struct {
	suite.Suite
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	userID            string
	err               error
}

func TestDeleteUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteUserUseCaseTestSuite))
}

func (suite *DeleteUserUseCaseTestSuite) SetupTest() {
	suite.mockUserService = new(service.MockUserService)
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.userID = "1"
}

func (suite *DeleteUserUseCaseTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("DeleteUser", suite.userID).Return(nil)
}

func (suite *DeleteUserUseCaseTestSuite) givenUserServiceReturnsError() {
	suite.mockUserService.On("DeleteUser", suite.userID).Return(user2.UserNotFound())
}

func (suite *DeleteUserUseCaseTestSuite) whenDeleteUserUseCaseIsCalled() {
	suite.err = user.DeleteUserUseCase(suite.userID, suite.mockUserService, suite.mockLoggerService)
}

func (suite *DeleteUserUseCaseTestSuite) thenExpectSuccess() {
	suite.NoError(suite.err)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *DeleteUserUseCaseTestSuite) thenExpectError() {
	suite.Error(suite.err)
	suite.Equal(suite.err.Error(), user2.UserNotFound().Error())
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *DeleteUserUseCaseTestSuite) TestDeleteUserUseCaseWithSuccessResult() {
	// Given
	suite.givenUserServiceReturnsSuccess()

	// When
	suite.whenDeleteUserUseCaseIsCalled()

	// Then
	suite.thenExpectSuccess()
}

func (suite *DeleteUserUseCaseTestSuite) TestDeleteUserUseCaseWithErrorResult() {
	// Given
	suite.givenUserServiceReturnsError()

	// When
	suite.whenDeleteUserUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}
