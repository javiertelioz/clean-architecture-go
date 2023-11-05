package auth

import (
	"errors"
	"testing"

	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/auth"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
	"github.com/stretchr/testify/suite"
)

type GetAccessTokenUseCaseTestSuite struct {
	suite.Suite
	mockUserService   *service.MockUserService
	mockCryptoService *service.MockCryptoService
	mockJwtService    *service.MockJwtService
	mockLoggerService *service.MockLoggerService
	user              *entity.User
	result            string
	err               error
}

func TestGetAccessTokenUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetAccessTokenUseCaseTestSuite))
}

func (suite *GetAccessTokenUseCaseTestSuite) SetupTest() {
	suite.mockUserService = new(service.MockUserService)
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.mockJwtService = new(service.MockJwtService)
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

func (suite *GetAccessTokenUseCaseTestSuite) givenUserServiceFindsUserByEmail(user *entity.User, err error) {
	suite.mockUserService.On("GetUserByEmail", suite.user.Email).Return(user, err)
}

func (suite *GetAccessTokenUseCaseTestSuite) givenCryptoServiceVerifiesPassword(err error) {
	suite.mockCryptoService.On("Verify", suite.user.Password, "password123").Return(err)
}

func (suite *GetAccessTokenUseCaseTestSuite) givenJWTServiceReturns(token string, err error) {
	suite.mockJwtService.On("Generate", suite.user).Return(token, err)
}

func (suite *GetAccessTokenUseCaseTestSuite) whenGetAccessTokenUseCaseIsCalled() {
	suite.result, suite.err = auth.GetAccessTokenUserUseCase(
		suite.user.Email,
		suite.user.Password,
		suite.mockCryptoService,
		suite.mockJwtService,
		suite.mockUserService,
		suite.mockLoggerService,
	)
}

func (suite *GetAccessTokenUseCaseTestSuite) thenExpectSuccess() {
	suite.NoError(suite.err)
	suite.NotNil(suite.result)
	suite.mockUserService.AssertExpectations(suite.T())
	suite.mockJwtService.AssertExpectations(suite.T())
}

func (suite *GetAccessTokenUseCaseTestSuite) thenExpectError() {
	suite.Error(suite.err)
	suite.Empty(suite.result)
	suite.mockUserService.AssertExpectations(suite.T())
	suite.mockJwtService.AssertExpectations(suite.T())
}

func (suite *GetAccessTokenUseCaseTestSuite) TestGetAccessTokenUseCaseWithSuccessResult() {
	// Given
	suite.givenUserServiceFindsUserByEmail(suite.user, nil)
	suite.givenCryptoServiceVerifiesPassword(nil)
	suite.givenJWTServiceReturns("token", nil)

	// When
	suite.whenGetAccessTokenUseCaseIsCalled()

	// Then
	suite.thenExpectSuccess()
}

func (suite *GetAccessTokenUseCaseTestSuite) TestGetAccessTokenUseCaseWithNoFoundUserResult() {
	// Given
	suite.givenUserServiceFindsUserByEmail(nil, exceptions.UserNotFound())

	// When
	suite.whenGetAccessTokenUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}

func (suite *GetAccessTokenUseCaseTestSuite) TestGetAccessTokenUseCaseWithBadCredentialsResult() {
	// Given
	suite.givenUserServiceFindsUserByEmail(suite.user, nil)
	suite.givenCryptoServiceVerifiesPassword(errors.New("password_wrong"))

	// When
	suite.whenGetAccessTokenUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}
