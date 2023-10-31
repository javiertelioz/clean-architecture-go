package auth

import (
	"github.com/google/uuid"
	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/auth"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type VerifyAccessTokenUseCaseTestSuite struct {
	suite.Suite
	mockJwtService    *service.MockJwtService
	mockLoggerService *service.MockLoggerService
	token             *entity.Token
	result            *entity.Token
	err               error
}

func TestVerifyAccessTokenUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(VerifyAccessTokenUseCaseTestSuite))
}

func (suite *VerifyAccessTokenUseCaseTestSuite) SetupTest() {
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.mockJwtService = new(service.MockJwtService)
	suite.token = &entity.Token{
		ID:        uuid.UUID([]byte("52fd4c33-2471-4e12-af95-a92dc1fc9d15")),
		UserID:    uint64(1),
		Role:      "admin",
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(1),
	}
}

func (suite *VerifyAccessTokenUseCaseTestSuite) givenJWTServiceReturns(token *entity.Token, err error) {
	suite.mockJwtService.On("Verify", "").Return(token, err)
}

func (suite *VerifyAccessTokenUseCaseTestSuite) whenVerifyAccessTokenUseCaseIsCalled() {
	suite.result, suite.err = auth.VerifyAccessTokenUserUseCase("", suite.mockJwtService, suite.mockLoggerService)
}

func (suite *VerifyAccessTokenUseCaseTestSuite) thenExpectSuccess() {
	suite.NoError(suite.err)
	suite.NotNil(suite.result)

	suite.mockJwtService.AssertExpectations(suite.T())
	suite.mockLoggerService.AssertExpectations(suite.T())
}

func (suite *VerifyAccessTokenUseCaseTestSuite) thenExpectError() {
	suite.Error(suite.err)
	suite.Nil(suite.result)

	suite.mockJwtService.AssertExpectations(suite.T())
	suite.mockLoggerService.AssertExpectations(suite.T())
}

func (suite *VerifyAccessTokenUseCaseTestSuite) TestVerifyAccessTokenUseCaseSuccessResult() {
	// Given
	suite.givenJWTServiceReturns(suite.token, nil)

	// When
	suite.whenVerifyAccessTokenUseCaseIsCalled()

	// Then
	suite.thenExpectSuccess()
}

func (suite *VerifyAccessTokenUseCaseTestSuite) TestVerifyAccessTokenUseCaseWithAuthInvalidTokenResult() {
	// Given
	suite.givenJWTServiceReturns(nil, exceptions.AuthInvalidToken())

	// When
	suite.whenVerifyAccessTokenUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}

func (suite *VerifyAccessTokenUseCaseTestSuite) TestVerifyAccessTokenUseCaseWithAuthExpiredTokenResult() {
	// Given
	suite.givenJWTServiceReturns(nil, exceptions.AuthExpiredToken())

	// When
	suite.whenVerifyAccessTokenUseCaseIsCalled()

	// Then
	suite.thenExpectError()
}
