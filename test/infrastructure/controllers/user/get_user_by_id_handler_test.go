package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/application/user_cases/user/mocks"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetUserByIdHandlerTestSuite struct {
	suite.Suite
	route             *gin.Engine
	controller        *controllers.UserController
	mockUserService   *mocks.MockUserService
	mockLoggerService *mocks.MockLoggerService
	mockCryptoService *mocks.MockCryptoService
	request           *http.Request
	response          *httptest.ResponseRecorder
	userId            string
	user              *entity.User
}

func TestGetUserByIdHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserByIdHandlerTestSuite))
}

func (suite *GetUserByIdHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.route = gin.Default()
	suite.mockUserService = new(mocks.MockUserService)
	suite.mockLoggerService = new(mocks.MockLoggerService)
	suite.mockCryptoService = new(mocks.MockCryptoService)
	suite.controller = controllers.NewUserController(suite.mockCryptoService, suite.mockUserService, suite.mockLoggerService)
	suite.user = &entity.User{
		ID:       1,
		LastName: "Doe",
		Name:     "John",
		Email:    "john@example.com",
		Phone:    "+123456789",
		Password: "password123",
	}
}

func (suite *GetUserByIdHandlerTestSuite) givenUserId(id string) {
	suite.userId = id
}

func (suite *GetUserByIdHandlerTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("GetUser", suite.userId).Return(suite.user, nil)
}

func (suite *GetUserByIdHandlerTestSuite) givenUserServiceReturnsError() {
	suite.mockUserService.On("GetUser", suite.userId).Return(nil, exceptions.UserNotFound())
}

func (suite *GetUserByIdHandlerTestSuite) whenCallGetUserByIdHandler() {
	suite.request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/user/%s", suite.userId), nil)
	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()

	suite.route.GET("/api/v1/user/:id", suite.controller.GetUserByIdHandler)
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *GetUserByIdHandlerTestSuite) thenReturnSuccessResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusOK, suite.response.Code)
	suite.Equal("Operation was successful", responseBody.Message)

	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUserByIdHandlerTestSuite) thenReturnNoFoundResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusNotFound, suite.response.Code)
	suite.Equal("USER_NOT_FOUND", responseBody.Message)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUserByIdHandlerTestSuite) TestGetUserByIdHandlerSuccess() {
	suite.givenUserId("1")
	suite.givenUserServiceReturnsSuccess()
	suite.whenCallGetUserByIdHandler()
	suite.thenReturnSuccessResponse()
}

func (suite *GetUserByIdHandlerTestSuite) TestGetUserByIdHandlerNotFound() {
	suite.givenUserId("10")
	suite.givenUserServiceReturnsError()
	suite.whenCallGetUserByIdHandler()
	suite.thenReturnNoFoundResponse()
}
