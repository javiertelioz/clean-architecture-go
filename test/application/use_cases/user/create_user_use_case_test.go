package user

import (
	"errors"
	"testing"

	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/user"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
	"github.com/stretchr/testify/suite"
)

type CreateUserUseCaseTestSuite struct {
	suite.Suite
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	mockCryptoService *service.MockCryptoService
	user              *entity.User
	result            *entity.User
	err               error
}

func TestCreateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUseCaseTestSuite))
}

func (suite *CreateUserUseCaseTestSuite) SetupTest() {
	suite.mockUserService = new(service.MockUserService)
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.mockCryptoService = new(service.MockCryptoService)
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

func (suite *CreateUserUseCaseTestSuite) givenUserServiceByEmailReturnsSuccess() {
	suite.mockUserService.On("GetUserByEmail", suite.user.Email).Return(suite.user, nil)
}

func (suite *CreateUserUseCaseTestSuite) givenUserServiceByEmailReturnsError() {
	suite.mockUserService.On("GetUserByEmail", suite.user.Email).Return(nil, exceptions.UserNotFound())
}

func (suite *CreateUserUseCaseTestSuite) givenCryptoServiceReturnsSuccess() {
	suite.mockCryptoService.On("Hash", suite.user.Password).Return("password", nil)
}

func (suite *CreateUserUseCaseTestSuite) givenCryptoServiceReturnsError() {
	suite.mockCryptoService.On("Hash", suite.user.Password).Return(nil, errors.New("password_wrong"))
}

func (suite *CreateUserUseCaseTestSuite) whenCreateUserUseCaseIsCalled() {
	suite.result, suite.err = user.CreateUserUseCase(suite.user, suite.mockCryptoService, suite.mockUserService, suite.mockLoggerService)
}

func (suite *CreateUserUseCaseTestSuite) thenExpectSuccess() {
	suite.NoError(suite.err)
	suite.NotNil(suite.result)
	suite.Equal(suite.user.Name, suite.result.Name)
	suite.mockUserService.AssertExpectations(suite.T())
	suite.mockCryptoService.AssertExpectations(suite.T())
}

func (suite *CreateUserUseCaseTestSuite) thenExpectError() {
	suite.Error(suite.err)
	suite.Nil(suite.result)
	//suite.mockCryptoService.AssertExpectations(suite.T())
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUserUseCaseWithSuccessResult() {
	// Given
	suite.givenCryptoServiceReturnsSuccess()
	suite.givenUserServiceByEmailReturnsError()
	suite.givenUserServiceReturnsSuccess()

	// When
	suite.whenCreateUserUseCaseIsCalled()

	// Then
	suite.thenExpectSuccess()
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUserUseCaseWithErrorResult() {
	// Given
	suite.givenCryptoServiceReturnsSuccess()
	suite.givenUserServiceByEmailReturnsError()
	suite.givenUserServiceReturnsError()

	// When
	suite.whenCreateUserUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}

func (suite *CreateUserUseCaseTestSuite) TestCreateUserUseCaseWithUserAlreadyExistResult() {
	// Given
	suite.givenCryptoServiceReturnsSuccess()
	suite.givenUserServiceByEmailReturnsSuccess()
	suite.givenUserServiceReturnsError()

	// When
	suite.whenCreateUserUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}
