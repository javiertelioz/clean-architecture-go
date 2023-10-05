package user

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/user"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/test/application/user_cases/user/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DeleteUserUseCaseTestSuite struct {
	suite.Suite
	mockUserService   *mocks.MockUserService
	mockLoggerService *mocks.MockLoggerService
	userID            string
	err               error
}

func TestDeleteUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteUserUseCaseTestSuite))
}

func (suite *DeleteUserUseCaseTestSuite) SetupTest() {
	suite.mockUserService = new(mocks.MockUserService)
	suite.mockLoggerService = new(mocks.MockLoggerService)
	suite.userID = "1"
}

func (suite *DeleteUserUseCaseTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("DeleteUser", suite.userID).Return(nil)
}

func (suite *DeleteUserUseCaseTestSuite) givenUserServiceReturnsError() {
	suite.mockUserService.On("DeleteUser", suite.userID).Return(exceptions.UserNotFound())
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
	suite.Equal(suite.err.Error(), exceptions.UserNotFound().Error())
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
