package auth

import (
	"errors"
	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/auth"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/test/application/user_cases/user/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GetAccessTokenUseCaseTestSuite struct {
	suite.Suite
	mockUserService   *mocks.MockUserService
	mockCryptoService *mocks.MockCryptoService
	mockJwtService    *mocks.MockJwtService
	mockLoggerService *mocks.MockLoggerService
	user              *entity.User
	result            string
	err               error
}

func TestGetAccessTokenUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetAccessTokenUseCaseTestSuite))
}

func (suite *GetAccessTokenUseCaseTestSuite) SetupTest() {
	suite.mockUserService = new(mocks.MockUserService)
	suite.mockLoggerService = new(mocks.MockLoggerService)
	suite.mockJwtService = new(mocks.MockJwtService)
	suite.mockCryptoService = new(mocks.MockCryptoService)
	suite.user = &entity.User{
		ID:       1,
		LastName: "Doe",
		Name:     "John",
		Email:    "john@example.com",
		Phone:    "+123456789",
		Password: "password123",
	}
}

func (suite *GetAccessTokenUseCaseTestSuite) givenUserServiceByEmailReturnsSuccess() {
	suite.mockUserService.On("GetUserByEmail", suite.user.Email).Return(suite.user, nil)
}

func (suite *GetAccessTokenUseCaseTestSuite) givenUserServiceByEmailReturnsError() {
	suite.mockUserService.On("GetUserByEmail", suite.user.Email).Return(nil, exceptions.UserNotFound())
}

func (suite *GetAccessTokenUseCaseTestSuite) givenCryptoServiceReturnsSuccess() {
	suite.mockCryptoService.On("Verify", suite.user.Password, "password123").Return(nil)
}

func (suite *GetAccessTokenUseCaseTestSuite) givenCryptoServiceReturnsError() {
	suite.mockCryptoService.On("Verify", suite.user.Password, "password123").Return(errors.New("password_wrong"))
}

func (suite *GetAccessTokenUseCaseTestSuite) givenJWTServiceReturnsSuccess() {
	suite.mockJwtService.On("Generate", suite.user).Return("token", nil)
}

func (suite *GetAccessTokenUseCaseTestSuite) givenJWTServiceReturnsError() {
	suite.mockJwtService.On("Generate", suite.user).Return(nil, exceptions.AuthExpiredToken())
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
	suite.givenUserServiceByEmailReturnsSuccess()
	suite.givenCryptoServiceReturnsSuccess()
	suite.givenJWTServiceReturnsSuccess()

	// When
	suite.whenGetAccessTokenUseCaseIsCalled()

	// Then
	suite.thenExpectSuccess()
}

func (suite *GetAccessTokenUseCaseTestSuite) TestGetAccessTokenUseCaseWithNoFoundUserResult() {
	// Given
	suite.givenUserServiceByEmailReturnsError()

	// When
	suite.whenGetAccessTokenUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}

func (suite *GetAccessTokenUseCaseTestSuite) TestGetAccessTokenUseCaseWithBadCredentialsResult() {
	// Given
	suite.givenUserServiceByEmailReturnsSuccess()
	suite.givenCryptoServiceReturnsError()

	// When
	suite.whenGetAccessTokenUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}
