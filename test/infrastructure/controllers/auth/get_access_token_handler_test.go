package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/dto"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetAccessTokenHandlerTestSuite struct {
	suite.Suite
	route             *gin.Engine
	controller        *controllers.AuthController
	mockCryptoService *service.MockCryptoService
	mockJwtService    *service.MockJwtService
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	user              *entity.User
	request           *http.Request
	response          *httptest.ResponseRecorder
	loginDTO          dto.LoginDTO
}

func TestGetAccessTokenHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(GetAccessTokenHandlerTestSuite))
}

func (suite *GetAccessTokenHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.route = gin.Default()
	suite.mockCryptoService = new(service.MockCryptoService)
	suite.mockJwtService = new(service.MockJwtService)
	suite.mockUserService = new(service.MockUserService)
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.controller = controllers.NewAuthController(
		suite.mockCryptoService,
		suite.mockJwtService,
		suite.mockUserService,
		suite.mockLoggerService)

	suite.loginDTO = dto.LoginDTO{
		Email:    "john@example.com",
		Password: "password123",
	}

	suite.user = &entity.User{
		ID:       1,
		LastName: "Doe",
		Name:     "John",
		Email:    "john@example.com",
		Phone:    "+123456789",
		Password: "password123",
	}
}

func (suite *GetAccessTokenHandlerTestSuite) givenUserServiceByEmailReturns(user *entity.User, err error) {
	suite.mockUserService.On("GetUserByEmail", suite.user.Email).Return(user, err)
}

func (suite *GetAccessTokenHandlerTestSuite) givenCryptoServiceReturns(err error) {
	suite.mockCryptoService.On("Verify", suite.user.Password, "password123").Return(err)
}

func (suite *GetAccessTokenHandlerTestSuite) givenJWTServiceReturnsGenerateToken(token string, err error) {
	suite.mockJwtService.On("Generate", suite.user).Return(token, err)
}

func (suite *GetAccessTokenHandlerTestSuite) whenCallGetAccessTokenHandler() {
	data, _ := json.Marshal(suite.loginDTO)
	suite.request, _ = http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(data))
	suite.response = httptest.NewRecorder()

	suite.route.POST("/api/v1/auth/login", suite.controller.GetAccessTokenHandler)
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *GetAccessTokenHandlerTestSuite) thenReturnSuccessResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusOK, suite.response.Code)
}

func (suite *GetAccessTokenHandlerTestSuite) thenReturnErrorResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusBadRequest, suite.response.Code)
	// suite.Equal("YOUR_ERROR_MESSAGE_HERE", responseBody.Message)
}

func (suite *GetAccessTokenHandlerTestSuite) TestGetAccessTokenHandlerSuccess() {
	suite.givenUserServiceByEmailReturns(suite.user, nil)
	suite.givenCryptoServiceReturns(nil)
	suite.givenJWTServiceReturnsGenerateToken("token", nil)

	suite.whenCallGetAccessTokenHandler()

	suite.thenReturnSuccessResponse()
}

/*func (suite *GetAccessTokenHandlerTestSuite) TestGetAccessTokenHandlerUserNotFound() {

	suite.givenUserServiceByEmailReturns(suite.user, nil)
	suite.givenCryptoServiceReturns(nil)

	suite.whenCallGetAccessTokenHandler()

	suite.thenReturnErrorResponse()
}*/

func (suite *GetAccessTokenHandlerTestSuite) TestGetAccessTokenHandlerWrongPassword() {
	suite.givenUserServiceByEmailReturns(suite.user, nil)
	suite.givenCryptoServiceReturns(errors.New("password_wrong"))

	suite.whenCallGetAccessTokenHandler()

	suite.thenReturnErrorResponse()
}

func (suite *GetAccessTokenHandlerTestSuite) TestGetAccessTokenHandlerJWTError() {
	suite.givenUserServiceByEmailReturns(suite.user, nil)
	suite.givenCryptoServiceReturns(nil)
	suite.givenJWTServiceReturnsGenerateToken("", exceptions.AuthExpiredToken())

	suite.whenCallGetAccessTokenHandler()

	suite.thenReturnErrorResponse()
}
